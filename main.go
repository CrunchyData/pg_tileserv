package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	// "text/template"
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
)

// programName is the name string we use
const programName string = "pg_tileserv"

// programVersion is the version string we use
const programVersion string = "0.1"

// worldMercWidth is the width of the Web Mercator plane
const worldMercWidth float64 = 40075016.6855784

// globalDb is a global database connection pointer
var globalDb *pgxpool.Pool = nil

/******************************************************************************/

func main() {

	viper.SetDefault("DbConnection", "sslmode=disable")
	viper.SetDefault("HttpHost", "0.0.0.0")
	viper.SetDefault("HttpPort", 7800)
	viper.SetDefault("UrlBase", "")
	viper.SetDefault("DefaultResolution", 4096)
	viper.SetDefault("DefaultBuffer", 256)
	viper.SetDefault("MaxFeaturesPerTile", 50)
	viper.SetDefault("DefaultMinZoom", 0)
	viper.SetDefault("DefaultMaxZoom", 22)
	viper.SetDefault("Debug", false)

	// Read environment configuration first
	if dbUrl := os.Getenv("DATABASE_URL"); dbUrl != "" {
		viper.Set("DbConnection", dbUrl)
	}

	// Read the commandline
	flagDebugOn := getopt.BoolLong("debug", 'd', "log debugging information")
	flagConfigFile := getopt.StringLong("config", 'c', "", "config file name")
	getopt.Parse()

	if *flagConfigFile != "" {
		viper.SetConfigFile(*flagConfigFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(fmt.Sprintf("/etc/%s", programName))
		viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", programName))
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debug(err)
		} else {
			log.Fatal(err)
		}
	}

	// Commandline over-rides config file for debugging
	if *flagDebugOn {
		viper.Set("Debug", true)
		log.SetLevel(log.TraceLevel)
	}

	// Report our status
	log.Infof("%s %s\n", programName, programVersion)
	log.Infof("Listening on: %s:%d", viper.GetString("HttpHost"), viper.GetInt("HttpPort"))

	// Load the global layer list right away
	// Also connects to database
	if err := LoadLayers(); err != nil {
		log.Fatal(err)
	}

	// Get to work
	handleRequests()
}

/******************************************************************************/

/******************************************************************************/

/*
func HandleRequestRoot(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"event": "handlerequest",
		"topic": "root",
	}).Trace("HandleRequestRoot")
	// Update the local copy
	// LoadLayerTableList()
	// LoadLayerFunctionList()

	type globalInfo struct {
		Tables    map[string]LayerTable
		Functions map[string]LayerFunction
	}
	info := globalInfo{
		globalLayerTables,
		globalLayerFunctions,
	}

	t, err := template.ParseFiles("assets/index.html")
	if err != nil {
		log.Warn(err)
	}
	// t.Execute(w, globalLayerTables)
	t.Execute(w, info)
}

func HandleRequestTablePreview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyrname := vars["name"]
	log.WithFields(log.Fields{
		"event": "handlerequest",
		"topic": "tablepreview",
		"key":   lyrname,
	}).Tracef("HandleRequestTablePreview: %s", lyrname)

	if lyr, ok := globalLayerTables[lyrname]; ok {
		t, err := template.ParseFiles("assets/preview.html")
		if err != nil {
			log.Warn(err)
		}
		t.Execute(w, lyr)
	}
}
*/

func requestListJson(w http.ResponseWriter, r *http.Request) error {
	log.WithFields(log.Fields{
		"event": "request",
		"topic": "layerlist",
	}).Trace("RequestLayerList")
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
	}).Tracef("RequestLayerDetail(%s)", lyrId)

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

	tilerequest := lyr.GetTileRequest(tile, r)
	mvt, errMvt := DBTileRequest(&tilerequest)
	if errMvt != nil {
		return errMvt
	}

	if _, errWrite := w.Write(mvt); errWrite != nil {
		return errWrite
	}
	w.Header().Add("Content-Type", "application/vnd.mapbox-vector-tile")

	return nil
}

type tileAppError struct {
	HttpCode int
	SrcErr   error
	Topic    string
	Message  string
}

func (tae tileAppError) Error() string {
	if tae.Message != "" {
		return fmt.Sprint("%s (%s)", tae.HttpCode, tae.Message, tae.SrcErr.Error())
	}
	return fmt.Sprint("%s", tae.HttpCode, tae.SrcErr.Error())
}

type tileAppHandler func(w http.ResponseWriter, r *http.Request) error

func (fn tileAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		if hdr, ok := r.Header["x-correlation-id"]; ok {
			log.WithField("correlation-id", hdr[0])
		}
		if e, ok := err.(tileAppError); ok {
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

// TODO, propogate id headers to error logging
// TODO, ensure all logging uses fields
// x-correlation-id

func handleRequests() {

	// creates a new instance of a mux router
	r := mux.NewRouter().StrictSlash(true)
	// r.HandleFunc("/", RequestListHtml).Methods("GET")
	// r.HandleFunc("/index.html", RequestListHtml).Methods("GET")
	// r.HandleFunc("/{name}.html", RequestDetailHtml).Methods("GET")
	r.Handle("/index.json", tileAppHandler(requestListJson))
	r.Handle("/{name}.json", tileAppHandler(requestDetailJson))
	r.Handle("/{name}/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", tileAppHandler(requestTile))

	// Allow CORS from anywhere
	corsOpt := handlers.AllowedOrigins([]string{"*"})

	// more "production friendly" timeouts
	// https://blog.simon-frey.eu/go-as-in-golang-standard-net-http-config-will-break-your-production/#You_should_at_least_do_this_The_easy_path
	s := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf("%s:%d", viper.GetString("HttpHost"), viper.GetInt("HttpPort")),
		Handler:      handlers.CORS(corsOpt)(r),
	}

	// TODO figure out how to gracefully shut down on ^C
	// and shut down all the database connections / statements
	log.Fatal(s.ListenAndServe())
}

/******************************************************************************/
