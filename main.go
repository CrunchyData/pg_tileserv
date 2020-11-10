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
	"strings"
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
)

// programName is the name string we use
const programName string = "pg_tileserv"

// programVersion is the version string we use
//const programVersion string = "0.1"
var programVersion string

// globalDb is a global database connection pointer
var globalDb *pgxpool.Pool = nil

// globalVersions holds the parsed output of postgis_full_version()
var globalVersions map[string]string = nil

// globalPostGISVersion is numeric, sortable postgis version (3.2.1 => 3002001)
var globalPostGISVersion int = 0

// serverBounds are the coordinate reference system and extent from
// which tiles are constructed
var globalServerBounds *Bounds = nil

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
	viper.SetDefault("MaxFeaturesPerTile", 10000)
	viper.SetDefault("DefaultMinZoom", 0)
	viper.SetDefault("DefaultMaxZoom", 22)
	viper.SetDefault("Debug", false)
	viper.SetDefault("AssetsPath", "./assets")
	// 1d, 1h, 1m, 1s, see https://golang.org/pkg/time/#ParseDuration
	viper.SetDefault("DbPoolMaxConnLifeTime", "1h")
	viper.SetDefault("DbPoolMaxConns", 4)
	viper.SetDefault("DbTimeout", 10)
	viper.SetDefault("CORSOrigins", []string{"*"})
	viper.SetDefault("BasePath", "/")

	viper.SetDefault("CoordinateSystem.SRID", 3857)
	// XMin, YMin, XMax, YMax, must be square
	viper.SetDefault("CoordinateSystem.Xmin", -20037508.3427892)
	viper.SetDefault("CoordinateSystem.Ymin", -20037508.3427892)
	viper.SetDefault("CoordinateSystem.Xmax", 20037508.3427892)
	viper.SetDefault("CoordinateSystem.Ymax", 20037508.3427892)
}

func main() {

	// Read the commandline
	flagDebugOn := getopt.BoolLong("debug", 'd', "log debugging information")
	flagConfigFile := getopt.StringLong("config", 'c', "", "full path to config file", "config.toml")
	flagHelpOn := getopt.BoolLong("help", 'h', "display help output")
	flagVersionOn := getopt.BoolLong("version", 'v', "display version number")
	getopt.Parse()

	if *flagHelpOn {
		getopt.PrintUsage(os.Stdout)
		os.Exit(1)
	}

	if *flagVersionOn {
		fmt.Printf("%s %s\n", programName, programVersion)
		os.Exit(0)
	}

	// Commandline over-rides config file for debugging
	if *flagDebugOn {
		viper.Set("Debug", true)
		log.SetLevel(log.TraceLevel)
	}

	if *flagConfigFile != "" {
		viper.SetConfigFile(*flagConfigFile)
	} else {
		viper.SetConfigName(programName)
		viper.AddConfigPath("./config")
		viper.AddConfigPath("/config")
		viper.AddConfigPath("/etc")
	}

	// Report our status
	log.Infof("%s %s", programName, programVersion)
	log.Info("Run with --help parameter for commandline options")

	// Read environment configuration first
	if dbUrl := os.Getenv("DATABASE_URL"); dbUrl != "" {
		viper.Set("DbConnection", dbUrl)
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

	basePath := viper.GetString("BasePath")
	log.Infof("Serving HTTP  at %s/", formatBaseURL(fmt.Sprintf("http://%s:%d",
		viper.GetString("HttpHost"), viper.GetInt("HttpPort")), basePath))
	log.Infof("Serving HTTPS at %s/", formatBaseURL(fmt.Sprintf("http://%s:%d",
		viper.GetString("HttpHost"), viper.GetInt("HttpsPort")), basePath))

	// Load the global layer list right away
	// Also connects to database
	if err := LoadLayers(); err != nil {
		log.Fatal(err)
	}

	// Read the postgis_full_version string and store
	// in a global for version testing
	if errv := LoadVersions(); errv != nil {
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
	lyrId := mux.Vars(r)["name"]
	log.WithFields(log.Fields{
		"event": "request",
		"topic": "layerpreview",
		"key":   lyrId,
	}).Tracef("requestPreview: %s", lyrId)

	// reqProperties := r.FormValue("properties")
	// reqLimit := r.FormValue("limit")
	// reqResolution := r.FormValue("resolution")
	// reqBuffer := r.FormValue("buffer")

	// Refresh the layers list
	if err := LoadLayers(); err != nil {
		return err
	}
	// Get the requested layer
	lyr, errLyr := GetLayer(lyrId)
	if errLyr != nil {
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

func requestListHtml(w http.ResponseWriter, r *http.Request) error {
	log.WithFields(log.Fields{
		"event": "request",
		"topic": "layerlist",
	}).Trace("requestListHtml")
	// Update the global in-memory list from
	// the database
	if err := LoadLayers(); err != nil {
		return err
	}
	jsonLayers := GetJsonLayers(r)

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

func requestListJson(w http.ResponseWriter, r *http.Request) error {
	log.WithFields(log.Fields{
		"event": "request",
		"topic": "layerlist",
	}).Trace("requestListJson")
	// Update the global in-memory list from
	// the database
	if err := LoadLayers(); err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	jsonLayers := GetJsonLayers(r)
	json.NewEncoder(w).Encode(jsonLayers)
	return nil
}

func requestDetailJson(w http.ResponseWriter, r *http.Request) error {
	lyrId := mux.Vars(r)["name"]
	log.WithFields(log.Fields{
		"event": "request",
		"topic": "layerdetail",
	}).Tracef("requestDetailJson(%s)", lyrId)

	// Refresh the layers list
	if err := LoadLayers(); err != nil {
		return err
	}

	lyr, errLyr := GetLayer(lyrId)
	if errLyr != nil {
		return errLyr
	}

	errWrite := lyr.WriteLayerJson(w, r)
	if errWrite != nil {
		return errWrite
	}
	return nil
}

func requestTile(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	lyr, errLyr := GetLayer(vars["name"])
	if errLyr != nil {
		return errLyr
	}
	tile, errTile := makeTile(vars)
	if errTile != nil {
		return errTile
	}

	log.WithFields(log.Fields{
		"event": "request",
		"topic": "tile",
		"key":   tile.String(),
	}).Tracef("RequestLayerTile: %s", tile.String())

	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("DbTimeout")*time.Second)
	defer cancel()

	tilerequest := lyr.GetTileRequest(tile, r)
	mvt, errMvt := DBTileRequest(ctx, &tilerequest)
	if errMvt != nil {
		return errMvt
	}

	w.Header().Add("Content-Type", "application/vnd.mapbox-vector-tile")

	if _, errWrite := w.Write(mvt); errWrite != nil {
		return errWrite
	}

	return nil
}

/******************************************************************************/

// tileAppError is an optional error structure functions can return
// if they want to specify the particular HTTP error code to be used
// in their error return
type tileAppError struct {
	HttpCode int
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
			if e.HttpCode == 0 {
				e.HttpCode = 500
			}
			if e.Topic != "" {
				log.WithField("topic", e.Topic)
			}
			log.WithField("key", e.Message)
			log.WithField("src", e.SrcErr.Error())
			log.Error(err)
			http.Error(w, e.Error(), e.HttpCode)
		} else {
			log.Error(err)
			http.Error(w, err.Error(), 500)
		}
	}
}

/******************************************************************************/

func SetSchemeHTTPS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Scheme = "https"
		next.ServeHTTP(w, r)
	})
}

/******************************************************************************/

func TileRouter() *mux.Router {
	// creates a new instance of a mux router
	r := mux.NewRouter().
		StrictSlash(true).
		PathPrefix(
			"/" +
				strings.TrimLeft(viper.GetString("BasePath"), "/"),
		).
		Subrouter()

	// Front page and layer list
	r.Handle("/", tileAppHandler(requestListHtml))
	r.Handle("/index.html", tileAppHandler(requestListHtml))
	r.Handle("/index.json", tileAppHandler(requestListJson))
	// Layer detail and demo pages
	r.Handle("/{name}.html", tileAppHandler(requestPreview))
	r.Handle("/{name}.json", tileAppHandler(requestDetailJson))
	// Tile requests
	r.Handle("/{name}/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", tileAppHandler(requestTile))
	return r
}

func handleRequests() {

	// Get a configured router
	r := TileRouter()

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
		Handler:      handlers.CompressHandler(handlers.CORS(corsOpt)(r)),
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
			Handler:      SetSchemeHTTPS(handlers.CompressHandler(handlers.CORS(corsOpt)(r))),
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
