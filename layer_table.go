package main

import (
	"fmt"

	// "github.com/lib/pq"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
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

type LayerTable struct {
	Id             string
	Schema         string
	Table          string
	Description    string
	Attributes     map[string]TableAttribute
	GeometryType   string
	IdColumn       string
	GeometryColumn string
	Srid           int
}

type TableAttribute struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	order       int
}

type TableDetailJson struct {
	Id           string           `json:"id"`
	Schema       string           `json:"schema"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Attributes   []TableAttribute `json:"attributes"`
	GeometryType string           `json:"geometrytype"`
	Center       [2]float64       `json:"center"`
	Bounds       [4]float64       `json:"bounds"`
	MinZoom      int              `json:"minzoom"`
	MaxZoom      int              `json:"maxzoom"`
	TileUrl      string           `json:"tileurl"`
	SourceLayer  string           `json:"sourcelayer"`
}

/********************************************************************************
 * Layer Interface
 */

func (lyr LayerTable) GetType() layerType {
	return layerTypeTable
}

func (lyr LayerTable) GetId() string {
	return lyr.Id
}

func (lyr LayerTable) GetDescription() string {
	return lyr.Description
}

func (lyr LayerTable) GetName() string {
	return lyr.Table
}

func (lyr LayerTable) GetSchema() string {
	return lyr.Schema
}

func (lyr LayerTable) GetTileRequest(tile Tile, req *http.Request) TileRequest {
	return TileRequest{} // TODO IMPLEMENT
}

func (lyr LayerTable) WriteLayerJson(w http.ResponseWriter, req *http.Request) error {
	jsonTableDetail, err := getTableDetailJson(&lyr, req)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonTableDetail)
	return nil // TODO IMPLEMENT
}

/********************************************************************************/

func getTableDetailJson(lyr *LayerTable, req *http.Request) (TableDetailJson, error) {
	td := TableDetailJson{
		Id:           lyr.Id,
		Schema:       lyr.Schema,
		Name:         lyr.Table,
		Description:  lyr.Description,
		GeometryType: lyr.GeometryType,
		MinZoom:      viper.GetInt("DefaultMinZoom"),
		MaxZoom:      viper.GetInt("DefaultMaxoom"),
		SourceLayer:  lyr.Id,
	}
	// Tile URL is relative to server base
	td.TileUrl = fmt.Sprintf("%s/%s/{z}/{x}/{y}.pbf", serverURLBase(req), lyr.Id)

	// Attributes:   lyr.Attributes, xxx

	// Read table bounds and convert to Json
	bnds, err := lyr.GetBounds()
	if err != nil {
		return td, err
	}
	td.Bounds[0] = bnds.Xmin
	td.Bounds[1] = bnds.Ymin
	td.Bounds[2] = bnds.Xmax
	td.Bounds[3] = bnds.Ymax
	td.Center[0] = (bnds.Xmin + bnds.Xmax) / 2.0
	td.Center[1] = (bnds.Ymin + bnds.Ymax) / 2.0
	return td, nil
}

func (lyr *LayerTable) GetBounds() (Bounds, error) {
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
	// queryBounds.Expand(float64(lyr.Buffer) / float64(lyr.Resolution))
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
	// mvtParams = append(mvtParams, fmt.Sprintf("%d", lyr.Resolution)) xxx
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
		// lyr.Resolution, xxxx
		// lyr.Buffer, xxx
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

func GetTableLayers() ([]LayerTable, error) {

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
			SELECT array_agg(ARRAY[sa.attname, st.typname, coalesce(da.description,''), sa,attnum::text]::text[] ORDER BY sa.attnum)
			FROM pg_attribute sa
			JOIN pg_type st ON sa.atttypid = st.oid
			LEFT JOIN pg_description da ON (c.oid = da.objoid and sa.attnum = da.objsubid)
			WHERE sa.attrelid = c.oid
			AND sa.attnum > 0
			AND NOT sa.attisdropped
			AND st.typname NOT IN ('geometry', 'geography')
		) AS props
	FROM pg_class c
	JOIN pg_namespace n ON (c.relnamespace = n.oid)
	JOIN pg_attribute a ON (a.attrelid = c.oid)
	JOIN pg_type t ON (a.atttypid = t.oid)
	LEFT JOIN pg_description d ON (c.oid = d.objoid and d.objsubid = 0)
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
		return nil, connerr
	}

	rows, err := db.Query(context.Background(), layerSql)
	if err != nil {
		return nil, connerr
	}

	// Reset array of layers
	layerTables := make([]LayerTable, 0)
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
			return nil, err
		}

		// We use https://godoc.org/github.com/jackc/pgtype#TextArray
		// here to scan the text[][] map of attribute name/type
		// created in the query. It gets a little ugly demapping the
		// pgx TextArray type, but it is at least native handling of
		// the array. It's complex because of PgSQL ARRAY generality
		// really, no fault of pgx
		attributes := make(map[string]TableAttribute)

		arrLen := atts.Dimensions[0].Length
		arrStart := atts.Dimensions[0].LowerBound - 1
		elmLen := atts.Dimensions[1].Length
		for i := arrStart; i < arrLen; i++ {
			pos := i * elmLen
			elmId := atts.Elements[pos].String
			elm := TableAttribute{
				Name:        elmId,
				Type:        atts.Elements[pos+1].String,
				Description: atts.Elements[pos+2].String,
			}
			elm.order, _ = strconv.Atoi(atts.Elements[pos+2].String)

			attributes[elmId] = elm
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
		}

		layerTables = append(layerTables, lyr)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return layerTables, nil
}
