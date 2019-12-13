package main

import (
	"fmt"
	// "github.com/lib/pq"
	"context"
	"strings"

	"github.com/jackc/pgtype"
	log "github.com/sirupsen/logrus"

	// Configuration
	"github.com/spf13/viper"
)

// x-correlation-id

// A Layer is a LayerTable or a LayerFunction

// in either case it should be able to generate
// SQL to produce tiles given an input tile

// type LayerTable struct {
type LayerTable struct {
	Id             string            `json:"id"`
	Schema         string            `json:"schema"`
	Table          string            `json:"table"`
	Description    string            `json:"description,omitempty"`
	Attributes     map[string]string `json:"attributes,omitempty"`
	GeometryType   string            `json:"geometry_type"`
	Center         []float64         `json:"center,omitempty"`
	MinZoom        int               `json:"minzoom,omitempty"`
	MaxZoom        int               `json:"maxzoom,omitempty"`
	Tiles          string            `json:"tiles,omitempty"`
	SourceLayer    int               `json:"source-layer,omitempty"`
	IdColumn       string
	GeometryColumn string
	Srid           int
	Buffer         int
	Resolution     int
	bounds         *Bounds
}

type LayerJson struct {
	Id string
}

type TileRequest struct {
	tile       Tile
	sql        string
	limit      int
	attributes []string
}

func (lyr *LayerTable) GetBounds() (Bounds, error) {
	if lyr.bounds != nil {
		return *lyr.bounds, nil
	}
	bounds := Bounds{}
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
		return bounds, err
	}

	err = db.QueryRow(context.Background(), extentSql).Scan(&bounds.Xmin, &bounds.Ymin, &bounds.Xmax, &bounds.Ymax)
	if err != nil {
		return bounds, err
	}

	log.Debug(bounds)
	return bounds, nil
}

func (lyr *LayerTable) TileSql(tile *Tile) string {

	// need both the exact tile boundary for clipping and an
	// expanded version for querying
	tileBounds := tile.Bounds()
	queryBounds := tile.Bounds()
	queryBounds.Expand(float64(lyr.Buffer) / float64(lyr.Resolution))
	tileSql := tileBounds.SQL()
	tileQuerySql := queryBounds.SQL()
	// convert the attribute name/type map into a SQL query for all
	// attributes
	// TODO, support attribute restriction in tile query
	attrNames := make([]string, 0)
	for k := range lyr.Attributes {
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
		viper.GetInt("MaxFeaturesPerTile"),
		strings.Join(mvtParams, ", "))

	log.Debug(sql)
	return sql
}

func (lyr *LayerTable) GetTile(tile *Tile) ([]byte, error) {

	db, err := DbConnect()
	if err != nil {
		log.Fatal(err)
	}

	tileSql := lyr.TileSql(tile)
	row := db.QueryRow(context.Background(), tileSql)
	var mvtTile []byte
	err = row.Scan(&mvtTile)
	if err != nil {
		log.Warn(err)
		return nil, err
	} else {
		return mvtTile, nil
	}
}

// https://github.com/mapbox/tilejson-spec/tree/master/2.0.1
type TileJson struct {
	TileJson    string      `json:"tilejson"`
	Name        string      `json:"name"`
	Data        string      `json:"data,omitempty"`
	Description string      `json:"description,omitempty"`
	Version     string      `json:"version"`
	Attribution string      `json:"attribution,omitempty"`
	Template    string      `json:"template,omitempty"`
	Legend      string      `json:"legend,omitempty"`
	Scheme      string      `json:"scheme"`
	Tiles       []string    `json:"tiles"`
	Grids       []string    `json:"grids,omitempty"`
	MinZoom     int         `json:"minzoom"`
	MaxZoom     int         `json:"maxzoom"`
	Bounds      []float64   `json:"bounds"`
	Center      []float64   `json:"center"`
	Id          string      `json:"id"`
	LayerConfig LayerConfig `json:"layerconfig"`
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

func (lyr *LayerTable) GetTileJson() (TileJson, error) {
	// initialize struct with known constants
	tileJson := TileJson{
		Version:  "1.0.0",
		TileJson: "2.0.1",
		MinZoom:  viper.GetInt("DefaultMinZoom"),
		MaxZoom:  viper.GetInt("DefaultMaxZoom"),
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
	tileJson.Tiles[0] = fmt.Sprintf("%s/%s/{z}/{x}/{y}.pbf", viper.GetString("UrlBase"), lyr.Id)
	tileJson.Id = lyr.Id
	tileJson.Attribution = viper.GetString("Attribution")

	bounds, err := lyr.GetBounds()
	if err != nil {
		return tileJson, err
	}
	tileJson.Bounds = make([]float64, 4)
	tileJson.Bounds[0] = bounds.Xmin
	tileJson.Bounds[1] = bounds.Ymin
	tileJson.Bounds[2] = bounds.Xmax
	tileJson.Bounds[3] = bounds.Ymax
	tileJson.Center = make([]float64, 2)
	tileJson.Center[0] = (bounds.Xmin + bounds.Xmax) / 2.0
	tileJson.Center[1] = (bounds.Ymin + bounds.Ymax) / 2.0

	tileJson.LayerConfig.Id = lyr.Id
	tileJson.LayerConfig.SourceLayer = lyr.Id
	tileJson.LayerConfig.Source.Type = "vector"
	tileJson.LayerConfig.Source.Tiles = make([]string, 1)
	tileJson.LayerConfig.Source.Tiles[0] = fmt.Sprintf("%s/%s/{z}/{x}/{y}.pbf", viper.GetString("UrlBase"), lyr.Id)
	tileJson.LayerConfig.Source.MinZoom = viper.GetInt("DefaultMinZoom")
	tileJson.LayerConfig.Source.MaxZoom = viper.GetInt("DefaultMaxZoom")

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
		log.Fatalf("unsupported geometry type %s", lyr.GeometryType)
	}

	tileJson.LayerConfig.Type = layerType

	log.Debug(tileJson)

	return tileJson, nil
}

func LoadLayerTableList() {

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
		LEFT JOIN pg_index i ON (c.oid = i.indrelid AND i.indisprimary AND i.indnatts = 1)
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
	globalLayerTables = make(map[string]LayerTable)
	for rows.Next() {

		var (
			schema, table, description, geometry_column string
			srid                                        int
			geometry_type, id_column                    string
			// props                                       [][]string
			atts pgtype.TextArray
		)

		err := rows.Scan(&schema, &table, &description, &geometry_column,
			&srid, &geometry_type, &id_column, &atts)
		if err != nil {
			log.Fatal(err)
		}

		// We use https://godoc.org/github.com/jackc/pgtype#TextArray
		// here to scan the text[][] map of attribute name/type
		// created in the query. It gets a little ugly demapping the
		// pgx TextArray type, but it is at least native handling of
		// the array. It's complex because of PgSQL ARRAY generality
		// really, no fault of pgx
		attributes := make(map[string]string)

		arrLen := atts.Dimensions[0].Length
		arrStart := atts.Dimensions[0].LowerBound - 1
		elmLen := atts.Dimensions[1].Length
		for i := arrStart; i < arrLen; i++ {
			elmPos := i * elmLen
			attributes[atts.Elements[elmPos].String] = atts.Elements[elmPos+1].String
		}

		// "schema.tablename" is our unique key for table layers
		id := fmt.Sprintf("%s.%s", schema, table)
		lyr := LayerTable{
			Id:             id,
			Schema:         schema,
			Table:          table,
			Description:    description,
			GeometryColumn: geometry_column,
			Srid:           srid,
			GeometryType:   geometry_type,
			IdColumn:       id_column,
			Attributes:     attributes,
			Resolution:     viper.GetInt("DefaultResolution"),
			Buffer:         viper.GetInt("DefaultBuffer"),
		}

		globalLayerTables[id] = lyr
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
	return
}
