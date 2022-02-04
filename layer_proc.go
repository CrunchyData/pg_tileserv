package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"

	// Logging
	log "github.com/sirupsen/logrus"

	// Configuration
	"github.com/spf13/viper"
)

// LayerFunction provides metadata about the function
type LayerFunction struct {
	ID          string
	Schema      string
	Function    string
	Description string
	Arguments   map[string]FunctionArgument
	MinZoom     int
	MaxZoom     int
	Tiles       string
}

// FunctionArgument provides the metadata and order
// of arguments in function call.
type FunctionArgument struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Default string `json:"default,omitempty"`
	order   int
}

// FunctionDetailJSON gives the output structure for
// the function.
type FunctionDetailJSON struct {
	ID          string             `json:"id"`
	Schema      string             `json:"schema"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Arguments   []FunctionArgument `json:"arguments"`
	MinZoom     int                `json:"minzoom"`
	MaxZoom     int                `json:"maxzoom"`
	TileURL     string             `json:"tileurl"`
}

/********************************************************************************
 * Layer Interface
 */

func (lyr LayerFunction) GetType() LayerType {
	return LayerTypeFunction
}

func (lyr LayerFunction) GetID() string {
	return lyr.ID
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

func (lyr LayerFunction) GetTileRequest(tile Tile, r *http.Request) TileRequest {

	procArgs := lyr.getFunctionArgs(r.URL.Query())
	sql, data, _ := lyr.requestSQL(tile, procArgs)

	tr := TileRequest{
		LayerID: lyr.ID,
		Tile:    tile,
		SQL:     sql,
		Args:    data,
	}
	return tr
}

func (lyr LayerFunction) WriteLayerJSON(w http.ResponseWriter, req *http.Request) error {
	jsonTableDetail, err := lyr.getFunctionDetailJSON(req)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonTableDetail)
	// all good, no error
	return nil
}

/********************************************************************************/

func (lyr *LayerFunction) requestSQL(tile Tile, args map[string]string) (string, []interface{}, error) {
	// Need ordered list of named parameters and values to
	// pass into the Query
	keys := []string{"x => $1", "y => $2", "z => $3"}
	vals := []interface{}{tile.X, tile.Y, tile.Zoom}
	i := 4
	for k, v := range args {
		keys = append(keys, fmt.Sprintf("%s => $%d", k, i))
		vals = append(vals, v)
		i++
	}

	// Build the SQL
	sql := fmt.Sprintf("SELECT \"%s\".\"%s\"(%s)", lyr.Schema, lyr.Function, strings.Join(keys, ", "))
	log.WithFields(log.Fields{
		"event": "tile",
		"topic": "sql",
		"key":   sql,
	}).Debugf("requestSql: %s", sql)

	return sql, vals, nil
}

func (lyr *LayerFunction) getFunctionArgs(vals url.Values) map[string]string {
	funcArgs := make(map[string]string)
	for k, v := range vals {
		if arg, ok := lyr.Arguments[k]; ok {
			funcArgs[arg.Name] = v[0]
		}
	}
	log.WithFields(log.Fields{
		"event":    "tile",
		"topic":    "args",
		"function": "getFunctionArgs",
		"return":   funcArgs,
	}).Debugf("getFunctionArgs => %s", funcArgs)
	return funcArgs
}

func (lyr *LayerFunction) getFunctionDetailJSON(req *http.Request) (FunctionDetailJSON, error) {

	td := FunctionDetailJSON{
		ID:          lyr.ID,
		Schema:      lyr.Schema,
		Name:        lyr.Function,
		Description: lyr.Description,
		Arguments:   make([]FunctionArgument, 0),
		MinZoom:     viper.GetInt("DefaultMinZoom"),
		MaxZoom:     viper.GetInt("DefaultMaxZoom"),
	}
	// TileURL is relative to server base
	td.TileURL = fmt.Sprintf("%s/%s/{z}/{x}/{y}.pbf", serverURLBase(req), url.PathEscape(lyr.ID))

	tmpMap := make(map[int]FunctionArgument)
	tmpKeys := make([]int, 0, len(lyr.Arguments))
	for _, v := range lyr.Arguments {
		tmpMap[v.order] = v
		tmpKeys = append(tmpKeys, v.order)
	}
	sort.Ints(tmpKeys)
	for _, v := range tmpKeys {
		td.Arguments = append(td.Arguments, tmpMap[v])
	}
	return td, nil
}

func getFunctionLayers() ([]LayerFunction, error) {

	// Valid functions **must** have signature of
	// function(z integer, x integer, y integer) returns bytea
	layerSQL := `
		SELECT
			Format('%s.%s', n.nspname, p.proname) AS id,
			n.nspname AS nspname,
			p.proname AS proname,
			coalesce(d.description, '') AS description,
			coalesce(p.proargnames, ARRAY[]::text[]) AS argnames,
			coalesce(string_to_array(oidvectortypes(p.proargtypes),', '), ARRAY[]::text[]) AS argtypes,
			coalesce(string_to_array(regexp_replace(pg_get_expr(p.proargdefaults, 0::Oid), '''([a-zA-Z0-9_\-\.]+)''::[a-z1-9]+', '\1'),', ', 'g'), ARRAY[]::text[]) AS argdefaults
		FROM pg_proc p
		JOIN pg_namespace n ON (p.pronamespace = n.oid)
		LEFT JOIN pg_description d ON (p.oid = d.objoid)
		WHERE p.proargtypes[0:2] = ARRAY[23::oid, 23::oid, 23::oid]
		AND p.proargnames[1:3] = ARRAY['z'::text, 'x'::text, 'y'::text]
		AND prorettype = 17
		AND has_schema_privilege(n.oid, 'usage')
		AND has_function_privilege(Format('%s.%s(%s)', quote_ident(n.nspname), quote_ident(p.proname), oidvectortypes(proargtypes)), 'execute')
		ORDER BY 1
		`

	db, connerr := dbConnect()
	if connerr != nil {
		return nil, connerr
	}

	rows, err := db.Query(context.Background(), layerSQL)
	if err != nil {
		return nil, err
	}

	// Reset array of layers
	layerFunctions := make([]LayerFunction, 0)
	for rows.Next() {

		var (
			id, schema, function, description string
			argnames, argtypes, argdefaults   []string
		)

		err := rows.Scan(&id, &schema, &function, &description, &argnames, &argtypes, &argdefaults)
		if err != nil {
			log.Fatal(err)
		}

		args := make(map[string]FunctionArgument)
		arglen := len(argnames)
		defstart := arglen - len(argdefaults)
		// First three arguments have to be z, x, y
		for i := 3; i < arglen; i++ {
			argdef := ""
			if i-defstart >= 0 {
				argdef = argdefaults[i-defstart]
			}
			args[argnames[i]] = FunctionArgument{
				order:   i - 3,
				Name:    argnames[i],
				Type:    argtypes[i],
				Default: parseArgDefault(argdef),
			}
		}

		lyr := LayerFunction{
			ID:          id,
			Schema:      schema,
			Function:    function,
			Description: description,
			Arguments:   args,
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

// parseArgDefault parses a default for an argument to a function-based
// tile layer. Most default arguments don't require special handling,
// but some are returned quoted with a type; e.g. a negative integer
// is `'-123'::integer` instead of `-123`.
func parseArgDefault(arg string) string {
	// check for a value in the value::type format
	sp := strings.Split(arg, "::")

	if len(sp) > 1 {
		// join back all but the last split parts.
		// this allows for the edge case of a double colon :: in text strings
		val := strings.Join(sp[:len(sp)-1], "::")

		// check for a value wrapped in single quotes and return the value
		// with them stripped.  If the value is not wrapped in quotes,
		// fall back to returning the value as is.
		if val[0] == '\'' && val[len(val)-1] == '\'' {
			return val[1 : len(val)-1]
		}
	}

	return arg
}
