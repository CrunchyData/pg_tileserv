package main

import (
	"fmt"

	"net/http"
	"net/url"
	"strings"

	// Database
	"context"

	log "github.com/sirupsen/logrus"
	// "github.com/jackc/pgtype"
	// "fmt"
	// "fmt"
	// "github.com/lib/pq"
	// "context"
	// "github.com/jackc/pgtype"
	// log "github.com/sirupsen/logrus"
	// "strings"
)

// type LayerTable struct {
type LayerFunction struct {
	Id            string    `json:"id"`
	Schema        string    `json:"schema"`
	Function      string    `json:"function"`
	Description   string    `json:"description,omitempty"`
	Arguments     []string  `json:"arguments,omitempty"`
	ArgumentTypes []string  `json:"argument_types,omitempty"`
	Center        []float64 `json:"center,omitempty"`
	MinZoom       int       `json:"minzoom,omitempty"`
	MaxZoom       int       `json:"maxzoom,omitempty"`
	Tiles         string    `json:"tiles,omitempty"`
	SourceLayer   string    `json:"source-layer,omitempty"`
}

/********************************************************************************
 * Layer Interface
 */

func (lyr LayerFunction) GetType() layerType {
	return layerTypeFunction
}

func (lyr LayerFunction) GetId() string {
	return lyr.Id
}

func (lyr LayerFunction) GetDescription() string {
	return lyr.Description
}

func (lyr LayerFunction) GetName() string {
	return lyr.Function
}

func (lyr LayerFunction) GetSchema() string {
	return lyr.Schema
}

func (lyr LayerFunction) GetTileRequest(tile Tile, req *http.Request) TileRequest {
	return TileRequest{} // TODO IMPLEMENT
}

func (lyr LayerFunction) WriteLayerJson(w http.ResponseWriter, req *http.Request) error {
	return nil // TODO IMPLEMENT
}

/********************************************************************************/

func (lyr *LayerFunction) GetLayerFunctionArgs(vals url.Values) map[string]string {
	funcArgs := make(map[string]string)
	for _, arg := range lyr.Arguments {
		if val, ok := vals[arg]; ok {
			funcArgs[arg] = val[0]
		}
	}
	return funcArgs
}

func (lyr *LayerFunction) GetTile(tile *Tile, args map[string]string) ([]byte, error) {

	db, err := DbConnect()
	if err != nil {
		log.Fatal(err)
	}

	// Need ordered list of named parameters and values to
	// pass into the Query
	keys := make([]string, 0)
	vals := make([]interface{}, 0)
	i := 1
	for k, v := range args {
		keys = append(keys, fmt.Sprintf("%s => $%d", k, i))
		switch k {
		case "x":
			vals = append(vals, tile.X)
		case "y":
			vals = append(vals, tile.Y)
		case "z":
			vals = append(vals, tile.Zoom)
		default:
			vals = append(vals, v)
		}
		i += 1
	}

	// Build the SQL
	sql := fmt.Sprintf("SELECT %s(%s)", lyr.Id, strings.Join(keys, ", "))
	log.WithFields(log.Fields{
		"event": "function.gettile",
		"topic": "sql",
		"key":   sql,
	}).Debugf("Func GetTile: %s", sql)

	row := db.QueryRow(context.Background(), sql, vals...)
	var mvtTile []byte
	err = row.Scan(&mvtTile)
	if err != nil {
		log.Warn(err)
		return nil, err
	} else {
		return mvtTile, nil
	}
}

func GetFunctionLayers() ([]LayerFunction, error) {

	// Valid functions **must** have signature of
	// function(z integer, x integer, y integer) returns bytea
	layerSql := `
		SELECT
		Format('%s.%s', n.nspname, p.proname) AS id,
		n.nspname,
		p.proname,
		coalesce(d.description, '') AS description,
		coalesce(p.proargnames, ARRAY[]::text[]) AS argnames,
		coalesce(string_to_array(oidvectortypes(p.proargtypes),', '), ARRAY[]::text[]) AS argtypes
		FROM pg_proc p
		JOIN pg_namespace n ON (p.pronamespace = n.oid)
		LEFT JOIN pg_description d ON (p.oid = d.objoid)
		WHERE p.proargtypes[0:2] = ARRAY[23::oid, 23::oid, 23::oid]
		AND p.proargnames[1:3] = ARRAY['z'::text, 'x'::text, 'y'::text]
		AND prorettype = 17
		AND has_function_privilege(Format('%s.%s(%s)', n.nspname, p.proname, oidvectortypes(proargtypes)), 'execute') ;
		`

	db, connerr := DbConnect()
	if connerr != nil {
		return nil, connerr
	}

	rows, err := db.Query(context.Background(), layerSql)
	if err != nil {
		return nil, err
	}

	// Reset array of layers
	layerFunctions := make([]LayerFunction, 0)
	for rows.Next() {

		var (
			id, schema, function, description string
			argnames, argtypes                []string
		)

		err := rows.Scan(&id, &schema, &function, &description, &argnames, &argtypes)
		if err != nil {
			log.Fatal(err)
		}

		lyr := LayerFunction{
			Id:            id,
			Schema:        schema,
			Function:      function,
			Description:   description,
			Arguments:     argnames[3:],
			ArgumentTypes: argtypes[3:],
		}

		layerFunctions = append(layerFunctions, lyr)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, err
	}
	rows.Close()
	return layerFunctions, nil
}

func LoadLayerFunctionList() {

	// Valid functions **must** have signature of
	// function(z integer, x integer, y integer) returns bytea
	layerSql := `
		SELECT
		Format('%s.%s', n.nspname, p.proname) AS id,
		n.nspname,
		p.proname,
		coalesce(d.description, '') AS description,
		coalesce(p.proargnames, ARRAY[]::text[]) AS argnames,
		coalesce(string_to_array(oidvectortypes(p.proargtypes),', '), ARRAY[]::text[]) AS argtypes
		FROM pg_proc p
		JOIN pg_namespace n ON (p.pronamespace = n.oid)
		LEFT JOIN pg_description d ON (p.oid = d.objoid)
		WHERE p.proargtypes[0:2] = ARRAY[23::oid, 23::oid, 23::oid]
		AND p.proargnames[1:3] = ARRAY['z'::text, 'x'::text, 'y'::text]
		AND prorettype = 17
		AND has_function_privilege(Format('%s.%s(%s)', n.nspname, p.proname, oidvectortypes(proargtypes)), 'execute') ;
		`

	db, connerr := DbConnect()
	if connerr != nil {
		log.Fatal(connerr)
	}

	rows, err := db.Query(context.Background(), layerSql)
	if err != nil {
		log.Fatal(err)
	}

	// Reset array of layers
	globalLayerFunctions = make(map[string]LayerFunction)
	for rows.Next() {

		var (
			id, schema, function, description string
			argnames, argtypes                []string
		)

		err := rows.Scan(&id, &schema, &function, &description, &argnames, &argtypes)
		if err != nil {
			log.Fatal(err)
		}

		lyr := LayerFunction{
			Id:            id,
			Schema:        schema,
			Function:      function,
			Description:   description,
			Arguments:     argnames[3:],
			ArgumentTypes: argtypes[3:],
		}

		globalLayerFunctions[id] = lyr
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
	return
}
