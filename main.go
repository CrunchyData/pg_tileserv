package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	// REST routing
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/theckman/httpforwarded"

	// Database connectivity
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
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

const programName string = "pg_tileserv"
const programVersion string = "0.1"

// A global array of Layer where the state is held for performance
// Refreshed when LoadLayerTableList is called
// Key is of the form: schemaname.tablename
var globalLayerTables map[string]Layer

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
	viper.SetDefault("MaxFeaturesPerTile", 50000)
	viper.SetDefault("DefaultMinZoom", 0)
	viper.SetDefault("DefaultMaxZoom", 22)
	viper.SetDefault("Debug", false)
	viper.SetDefault("Attribution", "")

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
	LoadLayerTableList()
	LoadLayerFunctionList()

	// Get to work
	HandleRequests()
}

/******************************************************************************/

func DbConnect() (*pgxpool.Pool, error) {
	if globalDb == nil {
		var err error
		var config *pgxpool.Config
		dbConnection := viper.GetString("DbConnection")
		config, err = pgxpool.ParseConfig(dbConnection)
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
		dbName := config.ConnConfig.Config.Database
		dbUser := config.ConnConfig.Config.User
		dbHost := config.ConnConfig.Config.Host
		log.Infof("Connected as '%s' to '%s' @ '%s'", dbUser, dbName, dbHost)

		return globalDb, err
	}
	return globalDb, nil
}

/******************************************************************************/

func serverURLBase(r *http.Request) string {

	// Use configuration file settings if we have them
	if viper.GetString("UrlBase") != "" {
		return viper.GetString("UrlBase")
	}

	// Preferred scheme
	ps := "http"
	// Preferred host:port
	ph := strings.TrimRight(r.Host, "/")
	// Preferred base path
	pb := "/"

	// Check for the IETF standard "Forwarded" header
	// for reverse proxy information
	xf := http.CanonicalHeaderKey("Forwarded");
	if f, ok := r.Header[xf]; ok {
		if fm, err := httpforwarded.Parse(f); err == nil {
			ph = fm["host"][0]
			ps = fm["proto"][0]
			return fmt.Sprintf("%v://%v%v", ps, ph, pb)
		}
	}

	// Check the X-Forwarded-Host and X-Forwarded-Proto
	// headers
	xfh := http.CanonicalHeaderKey("X-Forwarded-Host");
	if fh, ok := r.Header[xfh]; ok {
		ph = fh[0]
	}

	xfp := http.CanonicalHeaderKey("X-Forwarded-Proto");
	if fp, ok := r.Header[xfp]; ok {
		ps = fp[0]
	}

	return fmt.Sprintf("%v://%v%v", ps, ph, pb)
}


func HandleRequestRoot(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"event": "handlerequest",
		"topic": "root",
	}).Trace("HandleRequestRoot")
	// Update the local copy
	LoadLayerTableList()
	LoadLayerFunctionList()

	type globalInfo struct {
		Tables    map[string]Layer
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
	LoadLayerTableList()
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
	LoadLayerFunctionList()
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
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

func HandleRequestTableTile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyrname := vars["name"]
	if lyr, ok := globalLayerTables[lyrname]; ok {
		tile, _ := MakeTile(vars)

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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/vnd.mapbox-vector-tile")
		_, err = w.Write(pbf)
		return
	}

}

func HandleRequestFunctionTile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyrname := vars["name"]
	if lyr, ok := globalLayerFunctions[lyrname]; ok {
		tile, _ := MakeTile(vars)
		log.WithFields(log.Fields{
			"event": "handlerequest",
			"topic": "proctile",
			"key":   tile.String(),
		}).Tracef("HandleRequestFunctionTile: %s", tile.String())

		// Replace with SQL fun
		procArgs := lyr.GetLayerFunctionArgs(r.URL.Query())
		pbf, err := lyr.GetTile(&tile, procArgs)

		if err != nil {
			// TODO return a 500 or something
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/vnd.mapbox-vector-tile")
		_, err = w.Write(pbf)
		return
	}

}

func HandleRequests() {

	// creates a new instance of a mux router
	r := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	r.HandleFunc("/", HandleRequestRoot).Methods("GET")
	r.HandleFunc("/index.html", HandleRequestRoot).Methods("GET")
	r.HandleFunc("/index.json", HandleRequestTableList).Methods("GET")
	r.HandleFunc("/{name}.json", HandleRequestTable).Methods("GET")
	r.HandleFunc("/{name}.html", HandleRequestTablePreview).Methods("GET")
	r.HandleFunc("/{name}/tilejson.json", HandleRequestTableTileJSON).Methods("GET")
	r.HandleFunc("/{name}/{zoom:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", HandleRequestTableTile).Methods("GET")

	r.HandleFunc("/func/index.json", HandleRequestFunctionList).Methods("GET")
	r.HandleFunc("/func/{name}.json", HandleRequestFunction).Methods("GET")
	r.HandleFunc("/func/{name}/{zoom:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", HandleRequestFunctionTile)

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
