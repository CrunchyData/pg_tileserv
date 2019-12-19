package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"
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

// Age = 198
// Cats = [ "Cauchy", "Plato" ]
// Pi = 3.14
// Perfection = [ 6, 28, 496, 8128 ]
// DOB = 1987-07-05T05:45:00Z

// Then you can load it into your Go program with something like

// type Config struct {
//     Age int
//     Cats []string
//     Pi float64
//     Perfection []int
//     DOB time.Time
// }

// var conf Config
// if _, err := toml.DecodeFile("something.toml", &conf); err != nil {
//     // handle error
// }

// type Coordinate struct {
// 	x, y float64
// }

// programName is the name string we use
const programName string = "pg_tileserv"

// programVersion is the version string we use
const programVersion string = "0.1"

// worldMercWidth is the width of the Web Mercator plane
const worldMercWidth float64 = 40075016.6855784

// A global array of Layer where the state is held for performance
// Refreshed when LoadLayerTableList is called
// Key is of the form: schemaname.tablename
var globalLayerTables map[string]LayerTable

// A global array of LayerFunc where the state is held for performance
// Refreshed when LoadLayerTableList is called
// Key is of the form: schemaname.procname
var globalLayerFunctions map[string]LayerFunction

// A global database connection pointer
var globalDb *pgxpool.Pool = nil

/******************************************************************************/

func main() {

	viper.SetDefault("DbConnection", "sslmode=disable")
	viper.SetDefault("HttpHost", "0.0.0.0")
	viper.SetDefault("HttpPort", 7800)
	viper.SetDefault("UrlBase", "http://localhost:7800")
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
	HandleRequests()
}

/******************************************************************************/

/******************************************************************************/

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

func HandleRequestTableList(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"event": "handlerequest",
		"topic": "tablelist",
	}).Trace("HandleRequestTableList")
	// Update the local copy
	// LoadLayerTableList()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(globalLayerTables)
}

func HandleRequestFunctionList(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"event": "handlerequest",
		"topic": "proclist",
	}).Trace("HandleRequestFunctionList")
	// Update the local copy
	// LoadLayerFunctionList()
	// todo ERROR on db
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(globalLayerFunctions)
}

func HandleRequestTable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyrname := vars["name"]
	log.WithFields(log.Fields{
		"event": "handlerequest",
		"topic": "table",
		"key":   lyrname,
	}).Tracef("HandleRequestTable: %s", lyrname)

	if lyr, ok := globalLayerTables[lyrname]; ok {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lyr)
	}
	// todo ERROR
}

func HandleRequestFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyrname := vars["name"]
	log.WithFields(log.Fields{
		"event": "handlerequest",
		"topic": "proc",
		"key":   lyrname,
	}).Tracef("HandleRequestFunction: %s", lyrname)

	if lyr, ok := globalLayerFunctions[lyrname]; ok {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lyr)
	}
	// todo ERROR
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

func MakeTile(vars map[string]string) (Tile, error) {
	// Route restriction should ensure these are numbers
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(vars["y"])
	zoom, _ := strconv.Atoi(vars["zoom"])
	ext := vars["ext"]
	tile := Tile{Zoom: zoom, X: x, Y: y, Ext: ext}
	if !tile.IsValid() {
		return tile, errors.New(fmt.Sprintf("invalid tile address %s", tile.String()))
	}
	return tile, nil
}

// func HandleRequestFunctionTile(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	lyrname := vars["name"]
// 	if lyr, ok := globalLayerFunctions[lyrname]; ok {
// 		tile, _ := MakeTile(vars)
// 		log.WithFields(log.Fields{
// 			"event": "handlerequest",
// 			"topic": "proctile",
// 			"key":   tile.String(),
// 		}).Tracef("HandleRequestFunctionTile: %s", tile.String())

// 		// Replace with SQL fun
// 		procArgs := lyr.GetLayerFunctionArgs(r.URL.Query())
// 		pbf, err := lyr.GetTile(&tile, procArgs)

// 		if err != nil {
// 			// TODO return a 500 or something
// 		}
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Add("Content-Type", "application/vnd.mapbox-vector-tile")
// 		_, err = w.Write(pbf)
// 		return
// 	}

// }

func RequestListJson(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"event": "request",
		"topic": "layerlist",
	}).Trace("RequestLayerList")
	// Update the global in-memory list from
	// the database
	if err := LoadLayers(); err != nil {
		// return nil, err
		return
	}
	w.Header().Add("Content-Type", "application/json")
	jsonLayers := GetJsonLayers(r)
	json.NewEncoder(w).Encode(jsonLayers)
}

func RequestDetailJson(w http.ResponseWriter, r *http.Request) {
	lyrId := mux.Vars(r)["name"]
	log.WithFields(log.Fields{
		"event": "request",
		"topic": "layerdetail",
	}).Tracef("RequestLayerDetail(%s)", lyrId)

	if err := LoadLayers(); err != nil {
		// return nil, error
		return
	}

	lyr, errLyr := GetLayer(lyrId)
	if errLyr != nil {
		// return nil, errLyr
		return
	}

	errWrite := lyr.WriteLayerJson(w, r)
	if errWrite != nil {
		// return nil, errWrite
		return
	}
}

func RequestLayerTile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyr, errLyr := GetLayer(vars["name"])
	if errLyr != nil {
		// return nil, errLyr
		return
	}
	tile, errTile := MakeTile(vars)
	if errTile != nil {
		// return nil, errTile
		return
	}

	log.WithFields(log.Fields{
		"event": "request",
		"topic": "tile",
		"key":   tile.String(),
	}).Tracef("RequestLayerTile: %s", tile.String())

	tilerequest := lyr.GetTileRequest(tile, r)
	mvt, errMvt := DBTileRequest(&tilerequest)
	if errMvt != nil {
		// return nil, errMvt
		return
	}

	w.Header().Add("Content-Type", "application/vnd.mapbox-vector-tile")

	if _, err := w.Write(mvt); err != nil {
		// return nil, errWrite
		return
	}
	return
}

func HandleRequests() {

	// creates a new instance of a mux router
	r := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	// r.HandleFunc("/", HandleRequestRoot).Methods("GET")
	// r.HandleFunc("/index.html", HandleRequestRoot).Methods("GET")
	// r.HandleFunc("/index.json", HandleRequestTableList).Methods("GET")
	// r.HandleFunc("/{name}.html", HandleRequestTablePreview).Methods("GET")
	// r.HandleFunc("/{name}/tilejson.json", HandleRequestTableTileJSON).Methods("GET")
	// r.HandleFunc("/{name}/{zoom:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", HandleRequestTableTile).Methods("GET")

	// r.HandleFunc("/func/index.json", HandleRequestFunctionList).Methods("GET")
	// r.HandleFunc("/func/{name}.json", HandleRequestFunction).Methods("GET")
	// r.HandleFunc("/func/{name}/{zoom:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", HandleRequestFunctionTile)

	r.HandleFunc("/index.json", RequestListJson).Methods("GET")
	r.HandleFunc("/{name}.json", RequestDetailJson).Methods("GET")
	r.HandleFunc("/{name}/{zoom:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", RequestLayerTile).Methods("GET")

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

type StatusMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
