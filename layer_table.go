package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	// Database
	"github.com/jackc/pgtype"

	// Logging
	log "github.com/sirupsen/logrus"

	// Configuration
	"github.com/spf13/viper"
)

type LayerTable struct {
	Id             string
	Schema         string
	Table          string
	Description    string
	Properties     map[string]TableProperty
	GeometryType   string
	IdColumn       string
	GeometryColumn string
	Srid           int
}

type TableProperty struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	order       int
}

type TableDetailJson struct {
	Id           string          `json:"id"`
	Schema       string          `json:"schema"`
	Name         string          `json:"name"`
	Description  string          `json:"description,omitempty"`
	Properties   []TableProperty `json:"properties,omitempty"`
	GeometryType string          `json:"geometrytype,omitempty"`
	Center       [2]float64      `json:"center"`
	Bounds       [4]float64      `json:"bounds"`
	MinZoom      int             `json:"minzoom"`
	MaxZoom      int             `json:"maxzoom"`
	TileUrl      string          `json:"tileurl"`
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

func (lyr LayerTable) WriteLayerJson(w http.ResponseWriter, req *http.Request) error {
	jsonTableDetail, err := lyr.getTableDetailJson(req)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonTableDetail)
	// all good, no error
	return nil
}

func (lyr LayerTable) GetTileRequest(tile Tile, r *http.Request) TileRequest {
	rp := lyr.getQueryParameters(r.URL.Query())
	sql, _ := lyr.requestSql(&tile, &rp)

	tr := TileRequest{
		LayerId: lyr.Id,
		Tile:    tile,
		Sql:     sql,
		Args:    nil,
	}
	return tr
}

/********************************************************************************/

type queryParameters struct {
	Limit      int
	Properties []string
	Resolution int
	Buffer     int
}

// getRequestIntParameter ignores missing parameters and non-integer parameters,
// returning the "unknown integer" value for this case, which is -1
func getQueryIntParameter(q url.Values, param string) int {
	sParam, ok := q[param]
	if ok {
		iParam, err := strconv.Atoi(sParam[0])
		if err == nil {
			return iParam
		}
	}
	return -1
}

// getRequestPropertiesParameter compares the properties in the request
// with the properties in the table layer, and returns a slice of
// just those that occur in both, or a slice of all table properties
// if there is not query parameter, or no matches
func (lyr *LayerTable) getQueryPropertiesParameter(q url.Values) []string {
	sAtts, haveProperties := q["properties"]
	lyrAtts := (*lyr).Properties
	queryAtts := make([]string, 0, len(lyrAtts))
	haveIdColumn := false

	if haveProperties {
		aAtts := strings.Split(sAtts[0], ",")
		for _, att := range aAtts {
			decAtt, err := url.QueryUnescape(att)
			if err == nil {
				decAtt = strings.Trim(decAtt, " ")
				att, ok := lyrAtts[decAtt]
				if ok {
					if att.Name == lyr.IdColumn {
						haveIdColumn = true
					}
					queryAtts = append(queryAtts, att.Name)
				}
			}
		}
	}
	// No request parameter or no matches, so we want to
	// return all the properties in the table layer
	if len(queryAtts) == 0 {
		for _, v := range lyrAtts {
			queryAtts = append(queryAtts, v.Name)
		}
	}
	if (!haveIdColumn) && lyr.IdColumn != "" {
		queryAtts = append(queryAtts, lyr.IdColumn)
	}
	return queryAtts
}

// getRequestParameters reads user-settables parameters
// from the request URL, or uses the system defaults
// if the parameters are not set
func (lyr *LayerTable) getQueryParameters(q url.Values) queryParameters {
	rp := queryParameters{
		Limit:      getQueryIntParameter(q, "limit"),
		Resolution: getQueryIntParameter(q, "resolution"),
		Buffer:     getQueryIntParameter(q, "buffer"),
		Properties: lyr.getQueryPropertiesParameter(q),
	}
	if rp.Limit < 0 {
		rp.Limit = viper.GetInt("MaxFeaturesPerTile")
	}
	if rp.Resolution < 0 {
		rp.Resolution = viper.GetInt("DefaultResolution")
	}
	if rp.Buffer < 0 {
		rp.Buffer = viper.GetInt("DefaultBuffer")
	}
	return rp
}

/********************************************************************************/

func (lyr *LayerTable) getTableDetailJson(req *http.Request) (TableDetailJson, error) {
	td := TableDetailJson{
		Id:           lyr.Id,
		Schema:       lyr.Schema,
		Name:         lyr.Table,
		Description:  lyr.Description,
		GeometryType: lyr.GeometryType,
		MinZoom:      viper.GetInt("DefaultMinZoom"),
		MaxZoom:      viper.GetInt("DefaultMaxZoom"),
	}
	// TileURL is relative to server base
	td.TileUrl = fmt.Sprintf("%s/%s/{z}/{x}/{y}.pbf", serverURLBase(req), lyr.Id)

	// Want to add the properties to the Json representation
	// in table order, which is fiddly
	tmpMap := make(map[int]TableProperty)
	tmpKeys := make([]int, 0, len(lyr.Properties))
	for _, v := range lyr.Properties {
		tmpMap[v.order] = v
		tmpKeys = append(tmpKeys, v.order)
	}
	sort.Ints(tmpKeys)
	for _, v := range tmpKeys {
		td.Properties = append(td.Properties, tmpMap[v])
	}

	// Read table bounds and convert to Json
	// which prefers an array form
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

func (lyr *LayerTable) GetBoundsExact() (Bounds, error) {
	bounds := Bounds{}
	extentSql := fmt.Sprintf(`
	WITH ext AS (
		SELECT
			coalesce(
				ST_Transform(ST_SetSRID(ST_Extent("%s"), %d), 4326),
				ST_MakeEnvelope(-180, -90, 180, 90, 4326)
			) AS geom
		FROM "%s"."%s"
	)
	SELECT
		ST_XMin(ext.geom) AS xmin,
		ST_YMin(ext.geom) AS ymin,
		ST_XMax(ext.geom) AS xmax,
		ST_YMax(ext.geom) AS ymax
	FROM ext
	`, lyr.GeometryColumn, lyr.Srid, lyr.Schema, lyr.Table)

	db, err := DbConnect()
	if err != nil {
		return bounds, err
	}
	var (
		xmin pgtype.Float8
		xmax pgtype.Float8
		ymin pgtype.Float8
		ymax pgtype.Float8
	)
	err = db.QueryRow(context.Background(), extentSql).Scan(&xmin, &ymin, &xmax, &ymax)
	if err != nil {
		return bounds, tileAppError{
			SrcErr:  err,
			Message: "Unable to calculate table bounds",
		}
	}

	bounds.Xmin = xmin.Float
	bounds.Ymin = ymin.Float
	bounds.Xmax = xmax.Float
	bounds.Ymax = ymax.Float
	bounds.sanitize()
	return bounds, nil
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

	var (
		xmin pgtype.Float8
		xmax pgtype.Float8
		ymin pgtype.Float8
		ymax pgtype.Float8
	)
	err = db.QueryRow(context.Background(), extentSql).Scan(&xmin, &ymin, &xmax, &ymax)
	if err != nil {
		return bounds, tileAppError{
			SrcErr:  err,
			Message: "Unable to calculate table bounds",
		}
	}

	// Failed to get estimate? Get the exact bounds.
	if xmin.Status == pgtype.Null {
		warning := fmt.Sprintf("Estimated extent query failed, ANALYZE %s.%s", lyr.Schema, lyr.Table)
		log.WithFields(log.Fields{
			"event": "request",
			"topic": "detail",
			"key":   warning,
		}).Warn(warning)
		return lyr.GetBoundsExact()
	}

	bounds.Xmin = xmin.Float
	bounds.Ymin = ymin.Float
	bounds.Xmax = xmax.Float
	bounds.Ymax = ymax.Float
	bounds.sanitize()
	return bounds, nil
}

func (lyr *LayerTable) requestSql(tile *Tile, qp *queryParameters) (string, error) {

	type sqlParameters struct {
		TileSql        string
		QuerySql       string
		Resolution     int
		Buffer         int
		Properties     string
		MvtParams      string
		Limit          string
		Schema         string
		Table          string
		GeometryColumn string
		Srid           int
	}

	// need both the exact tile boundary for clipping and an
	// expanded version for querying
	tileBounds := tile.Bounds()
	queryBounds := tile.Bounds()
	queryBounds.Expand(tile.Width() * float64(qp.Buffer) / float64(qp.Resolution))
	tileSql := tileBounds.SQL()
	tileQuerySql := queryBounds.SQL()

	// preserve case and special characters in column names
	// of SQL query by double quoting names
	attrNames := make([]string, 0, len(qp.Properties))
	for _, a := range qp.Properties {
		attrNames = append(attrNames, fmt.Sprintf("\"%s\"", a))
	}

	// only specify MVT format parameters we have configured
	mvtParams := make([]string, 0)
	mvtParams = append(mvtParams, fmt.Sprintf("'%s', %d", lyr.Id, qp.Resolution))
	if lyr.GeometryColumn != "" {
		mvtParams = append(mvtParams, fmt.Sprintf("'%s'", lyr.GeometryColumn))
	}
	// The idColumn parameter is PostGIS3+ only
	if globalPostGISVersion >= 3000000 && lyr.IdColumn != "" {
		mvtParams = append(mvtParams, fmt.Sprintf("'%s'", lyr.IdColumn))
	}

	sp := sqlParameters{
		TileSql:        tileSql,
		QuerySql:       tileQuerySql,
		Resolution:     qp.Resolution,
		Buffer:         qp.Buffer,
		Properties:     strings.Join(attrNames, ", "),
		MvtParams:      strings.Join(mvtParams, ", "),
		Schema:         lyr.Schema,
		Table:          lyr.Table,
		GeometryColumn: lyr.GeometryColumn,
		Srid:           lyr.Srid,
	}

	if qp.Limit > 0 {
		sp.Limit = fmt.Sprintf("LIMIT %d", qp.Limit)
	}

	tmplSql := `
	SELECT ST_AsMVT(mvtgeom, {{ .MvtParams }}) FROM (
		SELECT ST_AsMVTGeom(
			ST_Transform(t."{{ .GeometryColumn }}", 3857),
			bounds.geom_clip,
			{{ .Resolution }},
			{{ .Buffer }}
		  ) AS "{{ .GeometryColumn }}"
		  {{ if .Properties }}
		  , {{ .Properties }}
		  {{ end }}
		FROM "{{ .Schema }}"."{{ .Table }}" t, (
			SELECT {{ .TileSql }}  AS geom_clip,
					{{ .QuerySql }} AS geom_query
			) bounds
		WHERE ST_Intersects(t."{{ .GeometryColumn }}",
							ST_Transform(bounds.geom_query, {{ .Srid }}))
		{{ .Limit }}
	) mvtgeom
	`

	sql, err := renderSqlTemplate("tableTileSql", tmplSql, sp)
	if err != nil {
		return "", err
	}
	return sql, err
}

func GetTableLayers() ([]LayerTable, error) {

	layerSql := `
	SELECT
		Format('%s.%s', n.nspname, c.relname) AS id,
		n.nspname AS schema,
		c.relname AS table,
		coalesce(d.description, '') AS description,
		a.attname AS geometry_column,
		postgis_typmod_srid(a.atttypmod) AS srid,
		trim(trailing 'ZM' from postgis_typmod_type(a.atttypmod)) AS geometry_type,
		coalesce(ia.attname, '') AS id_column,
		(
			SELECT array_agg(ARRAY[sa.attname, st.typname, coalesce(da.description,''), sa.attnum::text]::text[] ORDER BY sa.attnum)
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
	WHERE c.relkind IN ('r', 'v')
		AND t.typname = 'geometry'
		AND has_table_privilege(c.oid, 'select')
		AND has_schema_privilege(n.oid, 'usage')
		AND postgis_typmod_srid(a.atttypmod) > 0
	ORDER BY 1
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
			id, schema, table, description, geometry_column string
			srid                                            int
			geometry_type, id_column                        string
			atts                                            pgtype.TextArray
		)

		err := rows.Scan(&id, &schema, &table, &description, &geometry_column,
			&srid, &geometry_type, &id_column, &atts)
		if err != nil {
			return nil, err
		}

		// We use https://godoc.org/github.com/jackc/pgtype#TextArray
		// here to scan the text[][] map of property name/type
		// created in the query. It gets a little ugly demapping the
		// pgx TextArray type, but it is at least native handling of
		// the array. It's complex because of PgSQL ARRAY generality
		// really, no fault of pgx
		properties := make(map[string]TableProperty)

		if atts.Status == pgtype.Present {
			arrLen := atts.Dimensions[0].Length
			arrStart := atts.Dimensions[0].LowerBound - 1
			elmLen := atts.Dimensions[1].Length
			for i := arrStart; i < arrLen; i++ {
				pos := i * elmLen
				elmId := atts.Elements[pos].String
				elm := TableProperty{
					Name:        elmId,
					Type:        atts.Elements[pos+1].String,
					Description: atts.Elements[pos+2].String,
				}
				elm.order, _ = strconv.Atoi(atts.Elements[pos+3].String)
				properties[elmId] = elm
			}
		}

		// "schema.tablename" is our unique key for table layers
		lyr := LayerTable{
			Id:             id,
			Schema:         schema,
			Table:          table,
			Description:    description,
			GeometryColumn: geometry_column,
			Srid:           srid,
			GeometryType:   geometry_type,
			IdColumn:       id_column,
			Properties:     properties,
		}

		layerTables = append(layerTables, lyr)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return layerTables, nil
}
