package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
	"strconv"
	"time"
)

// type Coordinate struct {
// 	x, y float64
// }

type Config struct {
	ConnStr            string `json:"connstr"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	Addr               string `json:"addr"`
	Program            string `json:"program"`
	Version            string `json:"version"`
	DefaultResolution  int    `json:"default_resolution"`
	DefaultBuffer      int    `json:"default_buffer"`
	MaxFeaturesPerTile int    `json:"max_features_per_tile"`
	Attribution        string `json:"attribution"`
}

// A global variable for configuration parameters and defaults
var myConfig Config = Config{
	ConnStr:            "dbname=pramsey sslmode=disable",
	Host:               "localhost",
	Port:               7800,
	Addr:               "http://localhost:7800",
	Program:            "pg_tileserv",
	Version:            "0.1",
	DefaultBuffer:      256,
	DefaultResolution:  4094,
	MaxFeaturesPerTile: 10000,
}

// type LayerFunction struct {
// 	namespace string
// 	funcname string
// }

// A global array of Layer where the state is held for performance
// Refreshed when ReadLayerList is called
var Layers map[string]Layer

func HandleRequestRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleRequestRoot")
	b, err := ioutil.ReadFile("assets/index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(b))
}

func HandleRequestIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleRequestIndex")
	// Update the local copy
	ReadLayerList()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Layers)
}

func HandleRequestLayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyrname := vars["name"]
	log.Println("HandleRequestLayer: %s", lyrname)

	if lyr, ok := Layers[lyrname]; ok {
		tileJson, err := lyr.GetTileJson()
		if err == nil {
			json.NewEncoder(w).Encode(tileJson)
		}
	}
}

func HandleRequestTile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	lyrname := vars["name"]
	if lyr, ok := Layers[lyrname]; ok {
		x, _ := strconv.Atoi(vars["x"])
		y, _ := strconv.Atoi(vars["y"])
		zoom, _ := strconv.Atoi(vars["zoom"])
		ext := vars["ext"]
		log.Println("HandleRequestTile: %d/%d/%d.%s", zoom, x, y, ext)
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

// func trace() (string, int, string) {
//     pc, file, line, ok := runtime.Caller(1)
//     if !ok { return "?", 0, "?" }

//     fn := runtime.FuncForPC(pc)
//     return file, line, fn.Name()
// }

func HandleRequests() {

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", HandleRequestRoot)
	myRouter.HandleFunc("/index.json", HandleRequestIndex)
	myRouter.HandleFunc("/{name}.json", HandleRequestLayer)
	myRouter.HandleFunc("/{name}/{zoom:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", HandleRequestTile)

	// more "production friendly" timeouts
	// https://blog.simon-frey.eu/go-as-in-golang-standard-net-http-config-will-break-your-production/#You_should_at_least_do_this_The_easy_path
	s := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf("%s:%d", myConfig.Host, myConfig.Port),
		Handler:      myRouter,
	}

	// TODO figure out how to gracefully shut down on ^C
	// and shut down all the database connections / statements
	log.Fatal(s.ListenAndServe())
}

/******************************************************************************/

func main() {

	log.Printf("%s: %s\n", myConfig.Program, myConfig.Version)
	log.Printf("Listening on: %s", myConfig.Addr)

	// Load the layer list right away
	ReadLayerList()

	HandleRequests()
}
