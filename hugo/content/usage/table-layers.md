---
title: "Table Layers"
date:
draft: false
weight: 200
---

By default, `pg_tileserv` will provide access to **only** those spatial tables and views that:

* your database connection has `SELECT` privileges for;
* include a geometry column;
* declare a geometry type; and,
* declare an SRID (spatial reference ID).

For example:
```sql
CREATE TABLE mytable (
    geom Geometry(Polygon, 4326),
    pid text,
    address text
);
GRANT SELECT ON mytable TO myuser;
```

To restrict access to a certain set of tables, use database security principles:

* Create a role with limited privileges
* Only grant `SELECT` to that role for tables you want to publish
* Only grant `EXECUTE` to that role for functions you want to publish
* Connect `pg_tileserv` to the database using that role

## Table Layer Detail JSON

In the detail JSON, each layer declares information relevant to setting up a map interface for the layer.

```json
{
   "id" : "public.ne_50m_admin_0_countries",
   "geometrytype" : "MultiPolygon",
   "name" : "ne_50m_admin_0_countries",
   "description" : "Natural Earth countries",
   "schema" : "public",
   "bounds" : [
      -180,
      -89.9989318847656,
      180,
      83.599609375
   ],
   "center" : [
      0,
      -3.19966125488281
   ],
   "tileurl" : "http://localhost:7800/public.ne_50m_admin_0_countries/{z}/{x}/{y}.pbf",
   "properties" : [
      {
         "name" : "gid",
         "type" : "int4",
         "description" : ""
      },{
         "name" : "featurecla",
         "description" : "",
         "type" : "varchar"
      },{
         "description" : "",
         "type" : "varchar",
         "name" : "name"
      },{
         "type" : "varchar",
         "description" : "",
         "name" : "name_long"
      }
   ],
   "minzoom" : 0,
   "maxzoom" : 22
}
```
* `id`, `name`, and `schema` are the fully qualified, table, and schema name of the database table.
* `bounds` and `center` give the extent and middle of the data collection, in geographic coordinates. The order of coordinates in bounds is [minlon, minlat, maxlon, maxlat]. The order of coordinates in center is [lon, lat].
* `tileurl` is the standard substitution pattern URL consumed by map clients like [Mapbox GL JS](https://docs.mapbox.com/mapbox-gl-js/api/) and [OpenLayers](https://openlayers.org).
* `properties` is a list of columns in the table, with their data types and descriptions. The `description` field can be set using the `COMMENT` SQL command, for example:

```sql
COMMENT ON COLUMN ne_50m_admin_0_countries.name_long IS 'This is the long name';
```

## Table Tile Request Customization

Most developers will use the `tileurl` as is, but it's possible to add parameters to the URL to customize behaviour at run time:

* `limit` controls the number of features to write to a tile. The default is 50000.
* `resolution` controls the resolution of a tile. The default is 4096 units per side for a tile.
* `buffer` controls the size of the extra data buffer for a tile. The default is 256 units.
* `properties` is a comma-separated list of properties to include in the tile. For wide tables with large numbers of columns, this allows a slimmer tile to be composed.

For example:

    http://localhost:7800/public.ne_50m_admin_0_countries/{z}/{x}/{y}.pbf?limit=100000&properties=name,long_name

We recommend avoiding commas in property names. If necessary, you can [URL encode](https://en.wikipedia.org/wiki/Percent-encoding) the comma in the name string before composing the comma-separated string of all names.

* `filter` is an expression in [CQL](../cql) specifying what features are included in the tile
* `filter-crs` is the SRID of the coordinate reference system of any geometry literals in the CQL expression (the default is 4326)

## Multi-Layer Tile Requests

For more complex applications, multi-layer tiles can be useful to cut down on the amount of HTTP requests to pull in vector tiles. Doing this with `pg_tileserv` is easy, just add additional tables to your request. You can add as many tables as you like to your request, just separate them with a comma.

For example:

    http://localhost:7800/public.ne_50m_admin_0_countries,public.ne_50m_airports/{z}/{x}/{y}.pbf
