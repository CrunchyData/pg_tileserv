package main

import (
	"fmt"
	// "github.com/lib/pq"
	"context"
	"github.com/jackc/pgtype"
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
	GeometryType   string            `json:"geometry_type"`
	Properties     map[string]string `json:"properties,omitempty"`
	GeometryColumn string            `json:"geometry_column"`
	Srid           int               `json:"srid"`
	Id             string            `json:"id"`
	IdColumn       string            `json:"id_column,omitempty"`
	Resolution     int               `json:"resolution"`
	Buffer         int               `json:"buffer"`
	TileJson       TileJson          `json:"tilejson,omitempty"`
	LayerConfig    LayerConfig       `json:"layerconfig,omitempty"`
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
		globalConfig.MaxFeaturesPerTile,
		strings.Join(mvtParams, ", "))

	log.Println(sql)
	return sql
}

func (lyr *Layer) GetTile(tile *Tile) ([]byte, error) {

	db, err := DbConnect()
	if err != nil {
		log.Fatal(err)
	}

	tileSql := lyr.Sql(tile)
	rows, err := db.Query(context.Background(), tileSql)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var mvtTile []byte
	for rows.Next() {
		err = rows.Scan(&mvtTile)
		if err != nil {
			log.Println(err)
			rows.Close()
			return nil, err
		}
		// Check for errors from iterating over rows.
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		rows.Close()
		return nil, err
	}
	rows.Close()
	return mvtTile, nil
}

func (lyr *Layer) AddDetails() error {
	var err error
	if lyr.TileJson, err = lyr.GetTileJson(); err != nil {
		return err
	}
	if lyr.LayerConfig, err = lyr.GetLayerConfig(lyr.TileJson); err != nil {
		return err
	}
	return nil
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
	tileJson.Tiles[0] = fmt.Sprintf("%s/%s/{z}/{x}/{y}.pbf", globalConfig.UrlBase, lyr.Id)
	tileJson.Id = lyr.Id
	tileJson.Attribution = globalConfig.Attribution

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

	db, err := DbConnect()
	if err != nil {
		return tileJson, err
	}

	var (
		xmin, ymin, xmax, ymax float64
	)

	err = db.QueryRow(context.Background(), extentSql).Scan(&xmin, &ymin, &xmax, &ymax)
	if err != nil {
		return tileJson, err
	}
	return tileJson, nil
}

// https://github.com/mapbox/tilejson-spec/tree/master/2.0.1
type LayerConfig struct {
	Id          string `json:"id"`
	SourceLayer string `json:"source-layer"`
	Source      struct {
		Type    string   `json:"type"`
		Tiles   []string `json:"tiles"`
		MinZoom int      `json:"minzoom"`
		MaxZoom int      `json:"maxzoom"`
	} `json:"source"`
	Type string `json:"type"`
	// Paint map[string]interface{} `json:"paint"`
}

func (lyr *Layer) GetLayerConfig(tileJson TileJson) (LayerConfig, error) {

	layerConfig := LayerConfig{
		Id:          lyr.Id,
		SourceLayer: lyr.Id,
	}

	// 'id': json['id'],
	// 'source-layer': json['id'],
	// 'source': {
	//   'type': 'vector',
	//   'tiles': json['tiles'],
	//   'minzoom': json['minzoom'],
	//   'maxzoom': json['maxzoom']
	// },
	// 'type': 'circle',
	// 'paint': {
	//   'circle-radius': 2,
	//   'circle-color': '#007cbf'
	// }

	layerConfig.Source.Type = "vector"
	layerConfig.Source.Tiles = tileJson.Tiles
	layerConfig.Source.MinZoom = tileJson.MinZoom
	layerConfig.Source.MaxZoom = tileJson.MaxZoom

	var layerType string
	switch lyr.GeometryType {
	case "Point", "MultiPoint":
		layerType = "circle"
	case "LineString", "MultiLineString":
		layerType = "line"
	case "Polygon", "MultiPolygon":
		layerType = "line"
		// layerType = "fill"
	default:
		log.Fatal("unsupported geometry type %s", lyr.GeometryType)
	}

	layerConfig.Type = layerType

	return layerConfig, nil
}
//				SELECT array_agg(concat_ws(',', sa.attname, st.typname))

func GetLayerTableList() {

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
				SELECT array_agg(ARRAY[sa.attname, st.typname]::text[])
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

	db, connerr := DbConnect()
	if connerr != nil {
		log.Fatal(connerr)
	}

	rows, err := db.Query(context.Background(), layerSql)
	if err != nil {
		log.Fatal(err)
	}

	// Reset array of layers
	globalLayers = make(map[string]Layer)
	for rows.Next() {

		var (
			schema, table, description, geometry_column string
			srid                                        int
			geometry_type, id_column                    string
			// props                                       [][]string
			props                                    pgtype.TextArray
		)

		err := rows.Scan(&schema, &table, &description, &geometry_column,
			&srid, &geometry_type, &id_column, &props)
		if err != nil {
			log.Fatal(err)
		}

		// We use https://godoc.org/github.com/jackc/pgtype#TextArray
		// here to scan the text[][] map of attribute name/type
		// created in the query. It gets a little ugly demapping the
		// pgx TextArray type, but it is at least native handling of
		// the array. It's complex because of PgSQL ARRAY generality
		// really, no fault of pgx
		properties := make(map[string]string)

		arrLen := props.Dimensions[0].Length
		arrStart := props.Dimensions[0].LowerBound - 1
		elmLen := props.Dimensions[1].Length
		for i := arrStart; i < arrLen; i++ {
			elmPos := i * elmLen
			properties[props.Elements[elmPos].String] = props.Elements[elmPos + 1].String
		}

		lyr := Layer{
			Id:             fmt.Sprintf("%s.%s", schema, table),
			Schema:         schema,
			Table:          table,
			Description:    description,
			GeometryColumn: geometry_column,
			Srid:           srid,
			GeometryType:   geometry_type,
			IdColumn:       id_column,
			Properties:     properties,
			Resolution:     globalConfig.DefaultResolution,
			Buffer:         globalConfig.DefaultBuffer,
		}

		globalLayers[id] = lyr
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
	return
}
