package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	// REST routing
	"github.com/gorilla/mux"

	// Database connectivity
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"

	// Logging
	log "github.com/sirupsen/logrus"

	// Configuration
	"github.com/spf13/viper"
	"github.com/pborman/getopt/v2"
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



const programName string = "pg_tileserv"
const programVersion string = "0.1"

// A global array of Layer where the state is held for performance
// Refreshed when LoadLayerTableList is called
// Key is of the form: schemaname.tablename
var globalLayerTables map[string]Layer

// A global array of LayerProc where the state is held for performance
// Refreshed when LoadLayerTableList is called
// Key is of the form: schemaname.procname
var globalLayerProcs map[string]LayerProc

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
	viper.SetDefault("MaxFeaturesPerTile", 50000)
	viper.SetDefault("DefaultMinZoom", 0)
	viper.SetDefault("DefaultMaxZoom", 22)
	viper.SetDefault("Debug", false)

	// Read environment configuration first
	if dbUrl := os.Getenv("DATABASE_URL"); dbUrl != "" {
		viper.Set("DbConnection", dbUrl)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(fmt.Sprintf("/etc/%s", programName))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", programName))
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		    log.Debug(err)
	    } else {
    		log.Fatal(err)
	    }
	}

	// Read the commandline
	flagDebug := getopt.BoolLong("debug", 'd', "bool", "log debugging information")
	getopt.Parse()
	if *flagDebug {
		viper.Set("Debug", true)
	}

	if (viper.GetBool("Debug")) {
		log.SetLevel(log.TraceLevel)
	}

	// Report our status
	log.Infof("%s %s\n", programName, programVersion)
	log.Infof("Listening on: %s:%d", viper.GetString("HttpHost"),  viper.GetInt("HttpPort"))

	// Load the global layer list right away
	// Also connects to database
	LoadLayerTableList()

	// Get to work
	HandleRequests()
}

/******************************************************************************/

func DbConnect() (*pgxpool.Pool, error) {
	if globalDb == nil {
		var err error
		var config *pgxpool.Config
		config, err = pgxpool.ParseConfig(viper.GetString("DbConnection"))
		if err != nil {
			log.Fatal(err)
		}

		// Read current log level and use one less-fine level
		// below that
		config.ConnConfig.Logger = logrusadapter.NewLogger(log.New())
		levelString, _ := (log.GetLevel() - 1).MarshalText()
		pgxLevel, _ := pgx.LogLevelFromString(string(levelString))
		config.ConnConfig.LogLevel = pgxLevel

		// Connect!
		globalDb, err = pgxpool.ConnectConfig(context.Background(), config)
		if err != nil {
			log.Fatal(err)
		}
		// pgHost := config.ConnConfig.Config.Host
		// pgDatabase := config.ConnConfig.Config.Database
		// pgUser := config.ConnConfig.Config.User
		// pgPort := config.ConnConfig.Config.Port
		// log.Info(config.ConnConfig.Config)
		log.Infof("Connected to: %s\n", viper.GetString("DbConnection"))
		return globalDb, err
	}
	return globalDb, nil
}


/******************************************************************************/

func HandleRequestRoot(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"event": "handlerequest",
		"topic": "root",
	}).Trace("HandleRequestRoot")
	// Update the local copy
	LoadLayerTableList()
	LoadLayerProcList()

	t, err := template.ParseFiles("assets/index.html")
	if err != nil {
		log.Warn(err)
	}
	t.Execute(w, globalLayerTables)
}

func HandleRequestTableList(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"event": "handlerequest",
		"topic": "tablelist",
	}).Trace("HandleRequestTableList")
	// Update the local copy
	LoadLayerTableList()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(globalLayerTables)
}

func HandleRequestProcList(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"event": "handlerequest",
		"topic": "proclist",
	}).Trace("HandleRequestProcList")
	// Update the local copy
	LoadLayerProcList()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(globalLayerProcs)
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
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lyr)
	}
}

func HandleRequestTableTileJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyrname := vars["name"]
	log.WithFields(log.Fields{
		"event": "handlerequest",
		"topic": "tabletilejson",
		"key":   lyrname,
	}).Tracef("HandleRequestTableTileJSON: %s", lyrname)

	if lyr, ok := globalLayerTables[lyrname]; ok {
		tileJson, err := lyr.GetTileJson()
		log.Trace(tileJson)
		if err != nil {
			log.Warn(err)
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tileJson)
	}
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

func HandleRequestTableTile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	lyrname := vars["name"]
	if lyr, ok := globalLayerTables[lyrname]; ok {
		x, _ := strconv.Atoi(vars["x"])
		y, _ := strconv.Atoi(vars["y"])
		zoom, _ := strconv.Atoi(vars["zoom"])
		ext := vars["ext"]
		tile := Tile{Zoom: zoom, X: x, Y: y, Ext: ext}
		if !tile.IsValid() {
			log.Fatal("HandleRequestTableTile: invalid map tile")
		}
		log.WithFields(log.Fields{
			"event": "handlerequest",
			"topic": "tabletile",
			"key":   tile.String(),
		}).Tracef("HandleRequestTableTile: %s", tile.String())

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
	myRouter.HandleFunc("/index.json", HandleRequestTableList)
	myRouter.HandleFunc("/{name}.json", HandleRequestTable)
	myRouter.HandleFunc("/{name}.html", HandleRequestTablePreview)
	myRouter.HandleFunc("/{name}/tilejson.json", HandleRequestTableTileJSON)
	myRouter.HandleFunc("/{name}/{zoom:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", HandleRequestTableTile)

	myRouter.HandleFunc("/rpcs/index.json", HandleRequestProcList)
	// myRouter.HandleFunc("/rpcs/{name}.json", HandleRequestProc)
	// myRouter.HandleFunc("/rpcs/{name}/{zoom:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", HandleRequestProcTile)

	// more "production friendly" timeouts
	// https://blog.simon-frey.eu/go-as-in-golang-standard-net-http-config-will-break-your-production/#You_should_at_least_do_this_The_easy_path
	s := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf("%s:%d", viper.GetString("HttpHost"), viper.GetInt("HttpPort")),
		Handler:      myRouter,
	}

	// TODO figure out how to gracefully shut down on ^C
	// and shut down all the database connections / statements
	log.Fatal(s.ListenAndServe())
}

/******************************************************************************/

type Bounds struct {
	Minx float64 `json:"minx"`
	Miny float64 `json:"miny"`
	Maxx float64 `json:"maxx"`
	Maxy float64 `json:"maxx"`
}

func (b *Bounds) String() string {
	return fmt.Sprintf("{minx:%g, miny:%g, maxx:%g, maxy:%g}", b.Minx, b.Miny, b.Maxx, b.Maxy)
}

/******************************************************************************/

type StatusMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
