package main

import (
	// "bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
	"github.com/lib/pq"
	"strconv"
	"strings"
	"time"
)

// type Coordinate struct {
// 	x, y float64
// }

type Config struct {
	ConnStr           string `json:"connstr"`
	DefaultResolution int    `json:"default_resolution"`
	Host              string `json:"host"`
	Port              int    `json:"port"`
	Program           string `json:"program"`
	Version           string `json:"version"`
}

// A global variable for configuration parameters and defaults
var myConfig Config = Config{
	ConnStr: "dbname=pramsey sslmode=disable",
	Host:    "",
	Port:    1000,
	Program: "pg_tileserv",
	Version: "0.1",
}

// A Layer is a LayerTable or a LayerFunction
// in either case it should be able to generate
// SQL to produce tiles given an input tile

// type Layer interface {
// 	GetSQL(*Tile) string
// 	GetId() string
// }

type LayerList []Layer

// type LayerTable struct {
type Layer struct {
	Schema         string            `json:"schema"`
	Table          string            `json:"table"`
	Description    string            `json:"description,omitempty"`
	GeometryColumn string            `json:"geometry_column"`
	GeometryType   string            `json:"geometry_type"`
	Srid           int               `json:"srid"`
	Properties     map[string]string `json:"properties,omitempty"`
	Id             string            `json:"id"`
	IdColumn       string            `json:"id_column,omitempty"`
	Resolution     int               `json:"resolution"`
}

// type LayerFunction struct {
// 	namespace string
// 	funcname string
// }

// A global array of Layer where the state is held for performance
// Refreshed when ReadLayerList is called
var Layers map[string]Layer

func HandleRequestRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("Called: HandleRequestRoot")
	b, err := ioutil.ReadFile("assets/index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(b))
}

func HandleRequestIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleRequestIndex()")
	// Update the local copy
	ReadLayerList()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Layers)
}

func HandleRequestLayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lyrname := vars["name"]
	log.Println("HandleRequestLayer(%s)", lyrname)

	// Loop over all of our Layers
	// if the layer.Id equals the key we pass in
	// return the layer encoded as JSON
	if lyr, ok := Layers[lyrname]; ok {
		log.Println(lyr)
		json.NewEncoder(w).Encode(lyr)
	}
}

func HandleRequestTile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(vars["y"])
	zoom, _ := strconv.Atoi(vars["zoom"])
	ext := vars["ext"]
	log.Println("HandleRequestTile(%d/%d/%d.%s)", zoom, x, y, ext)
	tile := Tile{Zoom: zoom, X: x, Y: y, Ext: ext}
	if !tile.IsValid() {
		log.Fatal("HandleRequestTile: invalid map tile")
	}
	// Replace with SQL fun
	json.NewEncoder(w).Encode(tile)
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

// http://localhost:3000/public.geonames.jso
//
// {
//    "grids" : null,
//    "name" : "public.geonames",
//    "tilejson" : "2.2.0",
//    "data" : null,
//    "template" : null,
//    "scheme" : "xyz",
//    "version" : "1.0.0",
//    "center" : null,
//    "maxzoom" : 30,
//    "legend" : null,
//    "description" : null,
//    "tiles" : [
//       "http://localhost:3000/public.geonames/{z}/{x}/{y}.pbf"
//    ],
//    "bounds" : [
//       -180,
//       -90,
//       180,
//       90
//    ],
//    "id" : null,
//    "attribution" : null,
//    "minzoom" : 0
// }

//
// {
//    "public.geonames" : {
//       "resolution" : 0,
//       "id_column" : "id",
//       "geometry_type" : "Point",
//       "description" : "",
//       "properties" : {
//          "lon" : "float8",
//          "type" : "text",
//          "id" : "int4",
//          "ts" : "tsvector",
//          "lat" : "float8",
//          "state" : "text",
//          "name" : "text"
//       },
//       "schema" : "public",
//       "table" : "geonames",
//       "srid" : 4326,
//       "id" : "public.geonames",
//       "geometry_column" : "geom"
//    }
// }
//

func ReadLayerList() {

	layerSql := `
		SELECT
			n.nspname AS schema,
			c.relname AS table,
			coalesce(d.description, '') AS description,
			a.attname AS geometry_column,
			postgis_typmod_srid(a.atttypmod) AS srid,
			postgis_typmod_type(a.atttypmod) AS geometry_type,
			coalesce(ia.attname, '') AS id_column,
			(
				SELECT array_agg(concat_ws(',', sa.attname, st.typname))
				FROM pg_attribute sa
				JOIN pg_type st ON sa.atttypid = st.oid
				WHERE sa.attrelid = c.oid
				AND sa.attnum > 0
				AND NOT sa.attisdropped
				AND st.typname NOT IN ('geometry', 'geography')
			) AS props
		FROM pg_class c
		JOIN pg_namespace n ON (c.relnamespace = n.oid)
		JOIN pg_attribute a ON (a.attrelid = c.oid)
		JOIN pg_type t ON (a.atttypid = t.oid)
		LEFT JOIN pg_description d ON (c.oid = d.objoid)
		LEFT JOIN pg_index i ON (c.oid = i.indrelid AND i.indisprimary
		AND i.indnatts = 1)
		LEFT JOIN pg_attribute ia ON (ia.attrelid = i.indexrelid)
		LEFT JOIN pg_type it ON (ia.atttypid = it.oid AND it.typname in ('int2', 'int4', 'int8'))
		WHERE c.relkind = 'r'
		AND t.typname = 'geometry'
		AND has_table_privilege(c.oid, 'select')
		`

	db, err := sql.Open("postgres", myConfig.ConnStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(layerSql)
	if err != nil {
		log.Fatal(err)
	}

	// Reset array of layers
	Layers = make(map[string]Layer)
	for rows.Next() {

		var (
			schema, table, description, geometry_column string
			srid                                        int
			geometry_type, id_column                    string
			props                                       []string
		)

		err := rows.Scan(&schema, &table, &description, &geometry_column,
			&srid, &geometry_type, &id_column, pq.Array(&props))
		if err != nil {
			log.Fatal(err)
		}

		// we get back "name,type" from database query,
		// have to split them here, yuck, would rather have
		// a [][]string, but lib/pq doesn't support that yet
		properties := make(map[string]string)
		for _, att := range props {
			atts := strings.Split(att, ",")
			// TODO, guard against weird values
			properties[atts[0]] = atts[1]
		}

		// use fully qualified table name as id
		id := fmt.Sprintf("%s.%s", schema, table)

		lyr := Layer{
			Id:             id,
			Schema:         schema,
			Table:          table,
			Description:    description,
			GeometryColumn: geometry_column,
			Srid:           srid,
			GeometryType:   geometry_type,
			IdColumn:       id_column,
			Properties:     properties}

		Layers[id] = lyr
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

/******************************************************************************/

func main() {

	log.Printf("%s: %s\n", myConfig.Program, myConfig.Version)

	ReadLayerList()
	log.Println(Layers)

	HandleRequests()
}
