package main

import (
	"context"
	"log"
	// "fmt"
	// "fmt"
	// "github.com/lib/pq"
	// "context"
	// "github.com/jackc/pgtype"
	// log "github.com/sirupsen/logrus"
	// "strings"
)

// type LayerTable struct {
type LayerProc struct {
	Id            string   `json:"id"`
	Proc          string   `json:"proc"`
	Schema        string   `json:"schema"`
	Description   string   `json:"description,omitempty"`
	Arguments     []string `json:"arguments,omitempty"`
	ArgumentTypes []string `json:"argument_types,omitempty"`
}

func LoadLayerProcList() {

	// Valid procs **must** have signature of
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
	globalLayerProcs = make(map[string]LayerProc)
	for rows.Next() {

		var (
			id, schema, proc, description string
			argnames, argtypes            []string
		)

		err := rows.Scan(&id, &schema, &proc, &description, &argnames, &argtypes)
		if err != nil {
			log.Fatal(err)
		}

		lyr := LayerProc{
			Id:            id,
			Schema:        schema,
			Proc:          proc,
			Description:   description,
			Arguments:     argnames,
			ArgumentTypes: argtypes,
		}

		globalLayerProcs[id] = lyr
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
	return
}
