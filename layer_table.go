package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"strings"
)

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
	Buffer         int               `json:"buffer"`
}

func (lyr *Layer) Sql(tile *Tile) string {

	// need both the exact tile boundary for clipping and an
	// expanded version for querying
	tileBounds := tile.Bounds()
	tileSql := fmt.Sprintf("ST_MakeEnvelope(%g, %g, %g, %g, 3857)",
		tileBounds.Minx, tileBounds.Miny,
		tileBounds.Maxx, tileBounds.Maxy)
	tileQueryExpand := tile.Width() * float64(lyr.Buffer) / float64(lyr.Resolution)
	tileQuerySql := fmt.Sprintf("ST_Expand(%s, %g)", tileSql, tileQueryExpand)
	// convert the attribute name/type map into a SQL query for all
	// attributes
	// TODO, support attribute restriction in tile query
	attrNames := make([]string, 0)
	for k := range lyr.Properties {
		attrNames = append(attrNames, fmt.Sprintf("\"%s\"", k))
	}

	// only specify MVT format parameters we have configured
	mvtParams := make([]string, 0)
	mvtParams = append(mvtParams, fmt.Sprintf("'%s'::text", lyr.Id))
	mvtParams = append(mvtParams, fmt.Sprintf("%d", lyr.Resolution))
	if lyr.GeometryColumn != "" {
		mvtParams = append(mvtParams, fmt.Sprintf("'%s'::text", lyr.GeometryColumn))
	}
	if lyr.IdColumn != "" {
		mvtParams = append(mvtParams, fmt.Sprintf("'%s'::text", lyr.IdColumn))
	}

	sqlTmpl := `
		WITH
		bounds AS (
		  SELECT %s AS geom_query,
		         %s AS geom_clip
		),
		mvtgeom AS (
		  SELECT ST_AsMVTGeom(ST_Transform(t.%s, 3857), bounds.geom_clip, %d, %d) AS geom,
		       %s
		  FROM "%s"."%s" t, bounds
		  WHERE ST_Intersects(t.%s, ST_Transform(bounds.geom_query, %d))
		  LIMIT %d
		)
		SELECT ST_AsMVT(mvtgeom.*, %s) FROM mvtgeom
		`

	sql := fmt.Sprintf(sqlTmpl,
		tileQuerySql,
		tileSql,
		lyr.GeometryColumn,
		lyr.Resolution,
		lyr.Buffer,
		strings.Join(attrNames, ", "),
		lyr.Schema,
		lyr.Table,
		lyr.GeometryColumn,
		lyr.Srid,
		myConfig.MaxFeaturesPerTile,
		strings.Join(mvtParams, ", "))

	log.Println(sql)
	return sql
}

func (lyr *Layer) GetTile(tile *Tile) ([]byte, error) {

	db, err := sql.Open("postgres", myConfig.ConnStr)
	if err != nil {
		log.Fatal(err)
	}

	tileSql := lyr.Sql(tile)
	rows, err := db.Query(tileSql)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var mvtTile []byte
	rows.Next()
	err = rows.Scan(&mvtTile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return mvtTile, nil
}

// TODO, return the tile JSON information for this layer
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

// https://github.com/mapbox/tilejson-spec/tree/master/2.0.1
type TileJson struct {
	TileJson    string    `json:"tilejson"`
	Name        string    `json:"name"`
	Data        string    `json:"data,omitempty"`
	Description string    `json:"description,omitempty"`
	Version     string    `json:"version"`
	Attribution string    `json:"attribution,omitempty"`
	Template    string    `json:"template,omitempty"`
	Legend      string    `json:"legend,omitempty"`
	Scheme      string    `json:"scheme"`
	Tiles       []string  `json:"tiles"`
	Grids       []string  `json:"grids,omitempty"`
	MinZoom     int       `json:"minzoom"`
	MaxZoom     int       `json:"maxzoom"`
	Bounds      []float64 `json:"bounds"`
	Center      []float64 `json:"center"`
	Id          string    `json:"id"`
}

func (lyr *Layer) GetTileJson() (TileJson, error) {
	// initialize struct with known constants
	tileJson := TileJson{
		Version:  "1.0.0",
		TileJson: "2.0.1",
		MinZoom:  0,
		MaxZoom:  25,
		Scheme:   "xyz",
	}

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

	tileJson.Name = lyr.Id
	tileJson.Description = lyr.Description
	tileJson.Tiles = make([]string, 1)
	tileJson.Tiles[0] = fmt.Sprintf("%s/%s/{z}/{x}/{y}.pbf", myConfig.Addr, lyr.Id)
	tileJson.Id = lyr.Id
	tileJson.Attribution = myConfig.Attribution

	extentSql := fmt.Sprintf(`
		WITH ext AS (
			SELECT ST_Transform(ST_SetSRID(ST_EstimatedExtent('%s', '%s', '%s'), %d), 4326) AS geom
		)
		SELECT
			ST_XMin(ext.geom) AS xmin,
			ST_YMin(ext.geom) AS ymin,
			ST_XMax(ext.geom) AS xmax,
			ST_YMax(ext.geom) AS ymax
		FROM ext
		`, lyr.Schema, lyr.Table, lyr.GeometryColumn, lyr.Srid)

	db, err := sql.Open("postgres", myConfig.ConnStr)
	if err != nil {
		return tileJson, err
	}

	rows, err := db.Query(extentSql)
	if err != nil {
		return tileJson, err
	}

	// Reset array of layers
	for rows.Next() {

		var (
			xmin, ymin, xmax, ymax float64
		)
		err := rows.Scan(&xmin, &ymin, &xmax, &ymax)
		if err == nil {
			tileJson.Bounds = make([]float64, 4)
			tileJson.Bounds[0] = xmin
			tileJson.Bounds[1] = ymin
			tileJson.Bounds[2] = xmax
			tileJson.Bounds[3] = ymax
			tileJson.Center = make([]float64, 2)
			tileJson.Center[0] = (xmax + xmin) / 2.0
			tileJson.Center[1] = (ymax + ymin) / 2.0
		}
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return tileJson, err
	}
	return tileJson, nil
}

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
		AND postgis_typmod_srid(a.atttypmod) > 0
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
			Properties:     properties,
			Resolution:     myConfig.DefaultResolution,
			Buffer:         myConfig.DefaultBuffer,
		}

		Layers[id] = lyr
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
