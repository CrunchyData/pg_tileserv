package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	// REST routing
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	// Database connectivity
	"github.com/jackc/pgx/v4/pgxpool"

	// Logging
	log "github.com/sirupsen/logrus"

	// Configuration
	"github.com/pborman/getopt/v2"
	"github.com/spf13/viper"

	// Template functions
	"github.com/Masterminds/sprig/v3"

	// Prometheus metrics
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// programName is the name string we use
const programName string = "pg_tileserv"

// programVersion is the version string we use
// const programVersion string = "0.1"
var programVersion string

// globalDb is a global database connection pointer
var globalDb *pgxpool.Pool

// globalVersions holds the parsed output of postgis_full_version()
var globalVersions map[string]string

// globalPostGISVersion is numeric, sortable postgis version (3.2.1 => 3002001)
var globalPostGISVersion int

// serverBounds are the coordinate reference system and extent from
// which tiles are constructed
var globalServerBounds = make(map[int]*Bounds)
var globalDefaultCoordinateSystem int
var globalProjectionBoundsTableName string

// timeToLive is the Cache-Control timeout value that will be advertised
// in the response headers
var globalTimeToLive = -1

// A global array of Layer where the state is held for performance
// Refreshed when LoadLayerTableList is called
// Key is of the form: schemaname.tablename
var globalLayers map[string]Layer
var globalLayersMutex = &sync.Mutex{}

/******************************************************************************/

func init() {
	viper.SetDefault("DbConnection", "sslmode=disable")
	viper.SetDefault("HttpHost", "0.0.0.0")
	viper.SetDefault("HttpPort", 7800)
	viper.SetDefault("HttpsPort", 7801)
	viper.SetDefault("TlsServerCertificateFile", "")
	viper.SetDefault("TlsServerPrivateKeyFile", "")
	viper.SetDefault("UrlBase", "")
	viper.SetDefault("DefaultResolution", 4096)
	viper.SetDefault("DefaultBuffer", 256)
	viper.SetDefault("MaxFeaturesPerTile", 50000)
	viper.SetDefault("DefaultMinZoom", 0)
	viper.SetDefault("DefaultMaxZoom", 22)
	viper.SetDefault("Debug", false)
	viper.SetDefault("ShowPreview", true)
	viper.SetDefault("AssetsPath", "./assets")
	// 1d, 1h, 1m, 1s, see https://golang.org/pkg/time/#ParseDuration
	viper.SetDefault("DbPoolMaxConnLifeTime", "1h")
	viper.SetDefault("DbPoolMaxConns", 4)
	viper.SetDefault("DbTimeout", 10)
	viper.SetDefault("CORSOrigins", []string{"*"})
	viper.SetDefault("BasePath", "/")
	viper.SetDefault("CacheTTL", 0)          // cache timeout in seconds
	viper.SetDefault("EnableMetrics", false) // Prometheus metrics

	viper.SetDefault("DefaultCoordinateSystem", 3857)
	// XMin, YMin, XMax, YMax, must be square
	viper.SetDefault("CoordinateSystem.3857.Xmin", -20037508.3427892)
	viper.SetDefault("CoordinateSystem.3857.Ymin", -20037508.3427892)
	viper.SetDefault("CoordinateSystem.3857.Xmax", 20037508.3427892)
	viper.SetDefault("CoordinateSystem.3857.Ymax", 20037508.3427892)

	viper.SetDefault("HealthEndpoint", "/health")
}

func main() {

	// Read the commandline
	flagDebugOn := getopt.BoolLong("debug", 'd', "log debugging information")
	flagConfigFile := getopt.StringLong("config", 'c', "", "full path to config file", "config.toml")
	flagHelpOn := getopt.BoolLong("help", 'h', "display help output")
	flagVersionOn := getopt.BoolLong("version", 'v', "display version number")
	flagHidePreview := getopt.BoolLong("no-preview", 'n', "hide web interface")
	flagHealthEndpoint := getopt.StringLong("health", 'e', "", "desired path to health endpoint, e.g. \"/health\"")
	getopt.Parse()

	if *flagHelpOn {
		getopt.PrintUsage(os.Stdout)
		os.Exit(1)
	}

	if *flagVersionOn {
		fmt.Printf("%s %s\n", programName, programVersion)
		os.Exit(0)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("ts")

	// Enable debug mode if specified by commandline argument, regardless of what is in config file
	if *flagDebugOn {
		viper.Set("Debug", true)
		log.SetLevel(log.TraceLevel)
	}

	if *flagConfigFile != "" {
		viper.SetConfigFile(*flagConfigFile)
	} else {
		viper.SetConfigName(programName)
		viper.SetConfigType("toml")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("/config")
		viper.AddConfigPath("/etc")
	}

	if *flagHidePreview {
		viper.Set("ShowPreview", false)
	}

	if *flagHealthEndpoint != "" {
		viper.Set("HealthEndpoint", *flagHealthEndpoint)
	}

	// Report our status
	log.Infof("%s %s", programName, programVersion)
	log.Info("Run with --help parameter for commandline options")

	// Read environment configuration first
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		viper.Set("DbConnection", dbURL)
		log.Info("Using database connection info from environment variable DATABASE_URL")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debugf("viper.ConfigFileNotFoundError: %s", err)
		} else {
			if _, ok := err.(viper.UnsupportedConfigError); ok {
				log.Debugf("viper.UnsupportedConfigError: %s", err)
			} else {
				log.Fatalf("Configuration file error: %s", err)
			}
		}
	} else {
		// Really would like to log location of filename we found...
		// 	log.Infof("Reading configuration file %s", cf)
		if cf := viper.ConfigFileUsed(); cf != "" {
			log.Infof("Using config file: %s", cf)
		} else {
			log.Info("Config file: none found, using defaults")
		}
	}

	// enable debug mode if specified in config file, even if not specified by commandline argument
	if viper.GetBool("Debug") == true {
		log.SetLevel(log.TraceLevel)
	}

	basePath := viper.GetString("BasePath")
	log.Infof("Serving HTTP  at %s/", formatBaseURL(fmt.Sprintf("http://%s:%d",
		viper.GetString("HttpHost"), viper.GetInt("HttpPort")), basePath))
	log.Infof("Serving HTTPS at %s/", formatBaseURL(fmt.Sprintf("http://%s:%d",
		viper.GetString("HttpHost"), viper.GetInt("HttpsPort")), basePath))

	globalDefaultCoordinateSystem = viper.GetInt("DefaultCoordinateSystem")
	log.Infof("Default CoordinateSystem: %d", globalDefaultCoordinateSystem)

	globalProjectionBoundsTableName = viper.GetString("ProjectionBoundsTableName")

	// Load the global layer list right away
	// Also connects to database
	if err := loadLayers(); err != nil {
		log.Fatal(err)
	}

	// Read the postgis_full_version string and store
	// in a global for version testing
	if errv := loadVersions(); errv != nil {
		log.Fatal(errv)
	}
	log.WithFields(log.Fields{
		"event":       "connect",
		"topic":       "versions",
		"postgis":     globalVersions["POSTGIS"],
		"geos":        globalVersions["GEOS"],
		"pgsql":       globalVersions["PGSQL"],
		"libprotobuf": globalVersions["LIBPROTOBUF"],
	}).Debugf("Connected to PostGIS version %s\n", globalVersions["POSTGIS"])

	// Get to work
	handleRequests()
}

/******************************************************************************/

func requestPreview(w http.ResponseWriter, r *http.Request) error {
	lyrID := mux.Vars(r)["name"]
	log.WithFields(log.Fields{
		"event": "request",
		"topic": "layerpreview",
		"key":   lyrID,
	}).Tracef("requestPreview: %s", lyrID)

	// reqProperties := r.FormValue("properties")
	// reqLimit := r.FormValue("limit")
	// reqResolution := r.FormValue("resolution")
	// reqBuffer := r.FormValue("buffer")

	// Refresh the layers list
	if err := loadLayers(); err != nil {
		return err
	}
	// Get the requested layer
	lyr, err := getLayer(lyrID)
	if err != nil {
		errLyr := tileAppError{
			HTTPCode: 404,
			SrcErr:   err,
		}
		return errLyr
	}

	switch lyr.(type) {
	case LayerTable:
		tmpl, err := template.ParseFiles(fmt.Sprintf("%s/preview-table.html", viper.GetString("AssetsPath")))
		if err != nil {
			return err
		}
		l, _ := lyr.(LayerTable)
		tmpl.Execute(w, l)
	case LayerFunction:
		tmpl, err := template.ParseFiles(fmt.Sprintf("%s/preview-function.html", viper.GetString("AssetsPath")))
		if err != nil {
			return err
		}
		l, _ := lyr.(LayerFunction)
		tmpl.Execute(w, l)
	default:
		return errors.New("unknown layer type") // never get here
	}
	return nil
}

func requestListHTML(w http.ResponseWriter, r *http.Request) error {
	log.WithFields(log.Fields{
		"event": "request",
		"topic": "layerlist",
	}).Trace("requestListHtml")
	// Update the global in-memory list from
	// the database
	if err := loadLayers(); err != nil {
		return err
	}
	jsonLayers := getJSONLayers(r)

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/index.html", viper.GetString("AssetsPath")))

	if err != nil {
		return err
	}

	t, err := template.New("index").Funcs(sprig.FuncMap()).Parse(string(content))

	if err != nil {
		return err
	}
	t.Execute(w, jsonLayers)
	return nil
}

func requestListJSON(w http.ResponseWriter, r *http.Request) error {
	log.WithFields(log.Fields{
		"event": "request",
		"topic": "layerlist",
	}).Trace("requestListJSON")
	// Update the global in-memory list from
	// the database
	if err := loadLayers(); err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	jsonLayers := getJSONLayers(r)
	json.NewEncoder(w).Encode(jsonLayers)
	return nil
}

func requestDetailJSON(w http.ResponseWriter, r *http.Request) error {
	lyrID := mux.Vars(r)["name"]
	log.WithFields(log.Fields{
		"event": "request",
		"topic": "layerdetail",
	}).Tracef("requestDetailJSON(%s)", lyrID)

	// Refresh the layers list
	if err := loadLayers(); err != nil {
		return err
	}

	lyr, err := getLayer(lyrID)
	if err != nil {
		errLyr := tileAppError{
			HTTPCode: 404,
			SrcErr:   err,
		}
		return errLyr
	}

	errWrite := lyr.WriteLayerJSON(w, r)
	if errWrite != nil {
		return errWrite
	}
	return nil
}

// requestTile handles a tile request for a given layer
func requestTile(r *http.Request, source string, srid *int) ([]byte, error) {
	vars := mux.Vars(r)

	lyr, err := getLayer(source)
	if err != nil {
		errLyr := tileAppError{
			HTTPCode: 404,
			SrcErr:   err,
		}
		return nil, errLyr
	}

	tile, errTile := makeTile(vars, srid)
	if errTile != nil {
		return nil, errTile
	}

	log.WithFields(log.Fields{
		"event": "request",
		"topic": "tile",
		"key":   tile.String(),
	}).Tracef("requestTile: %s", tile.String())

	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("DbTimeout")*time.Second)
	defer cancel()

	tilerequest := lyr.GetTileRequest(tile, r)
	mvt, errMvt := dBTileRequest(ctx, &tilerequest)
	if errMvt != nil {
		return nil, errMvt
	}

	return mvt, nil
}

// requestTiles handles a tile request for a given layer, including multi layer tile requests
func requestTiles(w http.ResponseWriter, r *http.Request) error {
	var layers []byte
	vars := mux.Vars(r)

	var srid *int
	sridParam := r.URL.Query().Get("srid")
	sridInt, err := strconv.Atoi(sridParam)
	if err == nil {
		srid = &sridInt
	}

	sources := strings.Split(vars["name"], ",")
	var extant []string
	for _, source := range sources {
		if !slices.Contains(extant, source) {
			layer, err := requestTile(r, source, srid)
			if err != nil {
				return err
			}
			layers = append(layers, layer...)
			extant = append(extant, source)
		} else {
			log.Debugf("Skipping duplicate layer %s in request %s", source, sources)
		}
	}

	w.Header().Add("Content-Type", "application/vnd.mapbox-vector-tile")

	if _, errWrite := w.Write(layers); errWrite != nil {
		return errWrite
	}

	return nil
}

// A simple health check endpoint
func healthCheck(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 OK"))
	return nil
}

/******************************************************************************/

// tileAppError is an optional error structure functions can return
// if they want to specify the particular HTTP error code to be used
// in their error return
type tileAppError struct {
	HTTPCode int
	SrcErr   error
	Topic    string
	Message  string
}

// Error prints out a reasonable string format
func (tae tileAppError) Error() string {
	if tae.Message != "" {
		return fmt.Sprintf("%s\n%s", tae.Message, tae.SrcErr.Error())
	}
	return fmt.Sprintf("%s", tae.SrcErr.Error())
}

// tileAppHandler is a function handler that can replace the
// existing handler and provide richer error handling, see below and
// https://blog.golang.org/error-handling-and-go
type tileAppHandler func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP logs as much useful information as possible in
// a field format for potential Json logging streams
// as well as returning HTTP error response codes on failure
// so clients can see what is going on
// TODO: return JSON document body for the HTTP error
func (fn tileAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL,
	}).Infof("%s %s", r.Method, r.URL)
	if err := fn(w, r); err != nil {
		if hdr, ok := r.Header["x-correlation-id"]; ok {
			log.WithField("correlation-id", hdr[0])
		}
		if e, ok := err.(tileAppError); ok {
			if e.HTTPCode == 0 {
				e.HTTPCode = 500
			}
			if e.Topic != "" {
				log.WithField("topic", e.Topic)
			}
			log.WithField("key", e.Message)
			log.WithField("src", e.SrcErr.Error())
			log.Error(err)
			http.Error(w, e.Error(), e.HTTPCode)
		} else {
			log.Error(err)
			http.Error(w, err.Error(), 500)
		}
	}
}

/******************************************************************************/

func setCacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ttl := getTTL()
		if ttl > 0 {
			ccVal := fmt.Sprintf("max-age=%d", ttl)
			w.Header().Set("Cache-Control", ccVal)
		}
		next.ServeHTTP(w, r)
	})
}

/******************************************************************************/

func tileRouter() *mux.Router {
	// creates a new instance of a mux router
	r := mux.NewRouter().
		StrictSlash(true).
		PathPrefix(
			"/" +
				strings.TrimLeft(viper.GetString("BasePath"), "/"),
		).
		Subrouter()

	// Front page and layer list
	if viper.GetBool("ShowPreview") {
		r.Handle("/", tileAppHandler(requestListHTML))
		r.Handle("/index.html", tileAppHandler(requestListHTML))
		r.Handle("/index.json", tileAppHandler(requestListJSON))
		// Layer detail and demo pages
		r.Handle("/{name}.html", tileAppHandler(requestPreview))
		r.Handle("/{name}.json", tileAppHandler(requestDetailJSON))
	}
	// Tile requests
	r.Handle("/{name}/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", tileMetrics(tileAppHandler(requestTiles)))

	if viper.GetBool("EnableMetrics") {
		r.Handle("/metrics", promhttp.Handler())
	}

	r.Handle(viper.GetString("HealthEndpoint"), tileAppHandler(healthCheck)).Methods(http.MethodGet)
	return r
}

func handleRequests() {

	// Get a configured router
	r := tileRouter()

	// Allow CORS from anywhere
	corsOrigins := viper.GetStringSlice("CORSOrigins")
	corsOpt := handlers.AllowedOrigins(corsOrigins)

	// Set a writeTimeout for the http server.
	// This value is the application's DbTimeout config setting plus a
	// grace period. The additional time allows the application to gracefully
	// handle timeouts on its own, canceling outstanding database queries and
	// returning an error to the client, while keeping the http.Server
	// WriteTimeout as a fallback.
	writeTimeout := (time.Duration(viper.GetInt("DbTimeout") + 5)) * time.Second

	// more "production friendly" timeouts
	// https://blog.simon-frey.eu/go-as-in-golang-standard-net-http-config-will-break-your-production/#You_should_at_least_do_this_The_easy_path
	s := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: writeTimeout,
		Addr:         fmt.Sprintf("%s:%d", viper.GetString("HttpHost"), viper.GetInt("HttpPort")),
		Handler:      setCacheControl(handlers.CompressHandler(handlers.CORS(corsOpt)(r))),
	}

	// start http service
	go func() {
		// ListenAndServe returns http.ErrServerClosed when the server receives
		// a call to Shutdown(). Other errors are unexpected.
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	tlsServerCert := viper.GetString("TlsServerCertificateFile")
	tlsServerPrivKey := viper.GetString("TlsServerPrivateKeyFile")
	var stls *http.Server
	doServeTLS := false
	// Attempt to use HTTPS only if server certificate and private key files specified
	if tlsServerCert != "" && tlsServerPrivKey != "" {
		doServeTLS = true
	}

	log.Infof("Serving HTTP  at %s:%d", viper.GetString("HttpHost"), viper.GetInt("HttpPort"))

	if doServeTLS {
		log.Infof("Serving HTTPS at %s:%d", viper.GetString("HttpHost"), viper.GetInt("HttpsPort"))
		stls = &http.Server{
			ReadTimeout:  1 * time.Second,
			WriteTimeout: writeTimeout,
			Addr:         fmt.Sprintf("%s:%d", viper.GetString("HttpHost"), viper.GetInt("HttpsPort")),
			Handler:      setCacheControl(handlers.CompressHandler(handlers.CORS(corsOpt)(r))),
			TLSConfig: &tls.Config{
				MinVersion: tls.VersionTLS12, // Secure TLS versions only
			},
		}

		// start https service
		go func() {
			// ListenAndServe returns http.ErrServerClosed when the server receives
			// a call to Shutdown(). Other errors are unexpected.
			if err := stls.ListenAndServeTLS(tlsServerCert, tlsServerPrivKey); err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()
	}

	// wait here for interrupt signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// Interrupt signal received:  Start shutting down
	log.Infoln("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), writeTimeout)
	defer cancel()
	s.Shutdown(ctx)
	if doServeTLS {
		stls.Shutdown(ctx)
	}

	if globalDb != nil {
		log.Debugln("Closing DB connections")
		globalDb.Close()
	}
	log.Infoln("Server stopped.")
}

/******************************************************************************/
