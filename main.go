package main

import (
	// "bytes"
	// "database/sql"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	// "github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/log/logrusadapter"

	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	// _ "github.com/lib/pq"
	"html/template"
	"io/ioutil"
	"net/http"
	// "os"
	"github.com/BurntSushi/toml"
	"os"
	"strconv"
	"time"
	log "github.com/sirupsen/logrus"
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


type Config struct {
	DbConnection       string `json:"db_connection"`
	HttpHost           string `json:"http_host"`
	HttpPort           int    `json:"http_port"`
	UrlBase            string `json:"url_base"`
	DefaultResolution  int    `json:"default_resolution"`
	DefaultBuffer      int    `json:"default_buffer"`
	MaxFeaturesPerTile int    `json:"max_features_per_tile"`
	Attribution        string `json:"attribution"`
	DefaultMinZoom      int    `json:"default_minzoom"`
	DefaultMaxZoom      int    `json:"default_maxzoom"`
}

// A global variable for configuration parameters and defaults
// var globalConfig Config

// For un-provided values, use the defaults
var globalConfig Config = Config{
	DbConnection:       "sslmode=disable",
	HttpHost:           "0.0.0.0",
	HttpPort:           7800,
	UrlBase:            "http://localhost:7800",
	DefaultBuffer:      256,
	DefaultResolution:  4094,
	MaxFeaturesPerTile: 50000,
	DefaultMinZoom:     0,
	DefaultMaxZoom:     25,
}


const programName string = "pg_tileserv"
const programVersion string = "0.1"

// A global array of Layer where the state is held for performance
// Refreshed when GetLayerTableList is called
var globalLayers map[string]Layer

// A global database connection pointer
var globalDb *pgxpool.Pool = nil

// type LayerFunction struct {
// 	namespace string
// 	funcname string
// }

/******************************************************************************/

func main() {

	log.Infof("%s %s\n", programName, programVersion)

	// Read environment configuration first
	if dbUrl := os.Getenv("DATABASE_URL"); dbUrl != "" {
		globalConfig.DbConnection = dbUrl
	}

	// Attempt to read and parse command line configuration
	if len(os.Args) > 1 {
		configFile := os.Args[1]
		if _, err := os.Stat(configFile); err == nil {
			log.Infof("Reading configuration file: %s\n", configFile)
			if _, err := toml.DecodeFile(configFile, &globalConfig); err != nil {
				log.Fatal(err)
			}
		}
	}

	// Report our status
	log.Infof("Listening on: %s:%d", globalConfig.HttpHost, globalConfig.HttpPort)

	// Load the global layer list right away
	// Also connects to database
	GetLayerTableList()

	// Get to work
	HandleRequests()
}

/******************************************************************************/

func DbConnect() (*pgxpool.Pool, error) {
	if globalDb == nil {
		var err error
		var config *pgxpool.Config
		config, err = pgxpool.ParseConfig(globalConfig.DbConnection)
		if err != nil {
			log.Fatal(err)
		}
		config.ConnConfig.Logger = logrusadapter.NewLogger(log.New())
		config.ConnConfig.LogLevel = pgx.LogLevelWarn
		globalDb, err = pgxpool.ConnectConfig(context.Background(), config)
		if err != nil {
			log.Fatal(err)
		}
		// pgHost := config.ConnConfig.Config.Host
		// pgDatabase := config.ConnConfig.Config.Database
		// pgUser := config.ConnConfig.Config.User
		// pgPort := config.ConnConfig.Config.Port
		// log.Info(config.ConnConfig.Config)
		log.Infof("Connected to: %s\n", globalConfig.DbConnection)
		return globalDb, err
	}
	return globalDb, nil
}

/******************************************************************************/

func AssetFileAsString(assetPath string) (asset string) {
	b, err := ioutil.ReadFile(assetPath)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func HandleRequestRoot(w http.ResponseWriter, r *http.Request) {
	log.Trace("HandleRequestRoot")
	// html := AssetFileAsString("assets/index.html")
	// fmt.Fprintf(w, html)
	GetLayerTableList()

	t, err := template.ParseFiles("assets/index.html")
	if err != nil {
		log.Warn(err)
	}
	t.Execute(w, globalLayers)
}

func HandleRequestLayerList(w http.ResponseWriter, r *http.Request) {
	log.Trace("HandleRequestIndex")
	// Update the local copy
	GetLayerTableList()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(globalLayers)
}

func HandleRequestLayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyrname := vars["name"]
	log.Tracef("HandleRequestLayer: %s", lyrname)

	if lyr, ok := globalLayers[lyrname]; ok {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lyr)
	}
}

func HandleRequestLayerTileJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyrname := vars["name"]
	log.Tracef("HandleRequestLayerTileJSON: %s", lyrname)

	if lyr, ok := globalLayers[lyrname]; ok {
		tileJson, err := lyr.GetTileJson()
		log.Trace(tileJson)
		if err != nil {
			log.Warn(err)
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tileJson)
	}
}

func HandleRequestLayerPreview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyrname := vars["name"]
	log.Tracef("HandleRequestLayerPreview: %s", lyrname)

	if lyr, ok := globalLayers[lyrname]; ok {
		t, err := template.ParseFiles("assets/preview.html")
		if err != nil {
			log.Warn(err)
		}
		t.Execute(w, lyr)
	}
}

func HandleRequestTile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	lyrname := vars["name"]
	if lyr, ok := globalLayers[lyrname]; ok {
		x, _ := strconv.Atoi(vars["x"])
		y, _ := strconv.Atoi(vars["y"])
		zoom, _ := strconv.Atoi(vars["zoom"])
		ext := vars["ext"]
		log.Debugf("HandleRequestTile: %d/%d/%d.%s", zoom, x, y, ext)
		tile := Tile{Zoom: zoom, X: x, Y: y, Ext: ext}
		if !tile.IsValid() {
			log.Fatal("HandleRequestTile: invalid map tile")
		}
		// Replace with SQL fun
		pbf, err := lyr.GetTile(&tile)
		if err != nil {
			// TODO return a 500 or something
		}
		w.Header().Add("Content-Type", "application/vnd.mapbox-vector-tile")
		_, err = w.Write(pbf)
		return
	}

}

func HandleRequests() {

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", HandleRequestRoot)
	myRouter.HandleFunc("/index.html", HandleRequestRoot)
	myRouter.HandleFunc("/index.json", HandleRequestLayerList)
	myRouter.HandleFunc("/{name}.json", HandleRequestLayer)
	myRouter.HandleFunc("/{name}.html", HandleRequestLayerPreview)
	myRouter.HandleFunc("/{name}/tilejson.json", HandleRequestLayerTileJSON)
	myRouter.HandleFunc("/{name}/{zoom:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", HandleRequestTile)

	// more "production friendly" timeouts
	// https://blog.simon-frey.eu/go-as-in-golang-standard-net-http-config-will-break-your-production/#You_should_at_least_do_this_The_easy_path
	s := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf("%s:%d", globalConfig.HttpHost, globalConfig.HttpPort),
		Handler:      myRouter,
	}

	// TODO figure out how to gracefully shut down on ^C
	// and shut down all the database connections / statements
	log.Fatal(s.ListenAndServe())
}

/******************************************************************************/

type Bounds struct {
	Minx float64  `json:"minx"`
	Miny float64  `json:"miny"`
	Maxx float64  `json:"maxx"`
	Maxy float64  `json:"maxx"`
}

func (b *Bounds) String() string {
	return fmt.Sprintf("{minx:%g, miny:%g, maxx:%g, maxy:%g}", b.Minx, b.Miny, b.Maxx, b.Maxy)
}

/******************************************************************************/

type StatusMessage struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

