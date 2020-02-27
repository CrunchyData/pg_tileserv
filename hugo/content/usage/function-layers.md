---
title: "Function Layers"
date:
draft: false
weight: 300
---

## Function Layer Detail JSON

In the detail JSON, each function declares information relevant to setting up a map interface for the layer.

Since functions generate tiles dynamically, the system cannot auto-discover properties such as extent, or center. However, the custom parameters as well as defaults can be read from the function definition and exposed in the detail JSON.
```json
{
   "name" : "parcels_in_radius",
   "id" : "public.parcels_in_radius",
   "schema" : "public",
   "description" : "Given the click point (click_lon, click_lat) and radius, returns all the parcels in the radius, clipped to the radius circle.",
   "minzoom" : 0,
   "arguments" : [
      {
         "default" : "-123.13",
         "name" : "click_lon",
         "type" : "double precision"
      },
      {
         "default" : "49.25",
         "name" : "click_lat",
         "type" : "double precision"
      },
      {
         "default" : "500.0",
         "type" : "double precision",
         "name" : "radius"
      }
   ],
   "maxzoom" : 22,
   "tileurl" : "http://localhost:7800/public.parcels_in_radius/{z}/{x}/{y}.pbf"
}
```
* `description` can be set using the `COMMENT ON FUNCTION` SQL command.
* `id`, `schema`, and `name` are the fully qualified name, schema, and function name, respectively.
* `minzoom` and `maxzoom` are the defaults as set in the configuration file.
* `arguments` is a list of argument names, with the data type and default value.

## Function Layer Examples

### Filtering example

This simple example returns a filtered subset of a table ([ne_50m_admin_0_countries](https://www.naturalearthdata.com/http//www.naturalearthdata.com/download/50m/cultural/ne_50m_admin_0_countries.zip) [EPSG:4326](https://epsg.io/4326)). The filter in this case is the first letter of the name.

Note that the `name_prefix` parameter includes a **default value**: this is useful for clients (like the preview interface for this server) that read arbitrary function definitions and need a default value to fill into interface fields.

This example also uses `ST_TileEnvelope()`, a utility function only available in PostGIS 3.0 and higher. See the notes below for a workaround using custom functions.
```sql
CREATE OR REPLACE
FUNCTION public.countries_name(
            z integer, x integer, y integer,
            name_prefix text default 'B')
RETURNS bytea
AS $$
    WITH
    bounds AS (
      SELECT ST_TileEnvelope(z, x, y) AS geom
    ),
    mvtgeom AS (
      SELECT ST_AsMVTGeom(ST_Transform(t.geom, 3857), bounds.geom) AS geom,
        t.name
      FROM ne_50m_admin_0_countries t, bounds
      WHERE ST_Intersects(t.geom, ST_Transform(bounds.geom, 4326))
      AND upper(t.name) LIKE (upper(name_prefix) || '%')
    )
    SELECT ST_AsMVT(mvtgeom, 'public.countries_name') FROM mvtgeom;
$$
LANGUAGE 'sql'
STABLE
PARALLEL SAFE;

COMMENT ON FUNCTION public.countries_name IS 'Filters the countries table by the initial letters of the name using the "name_prefix" parameter.';
```
Some notes about this function:

* The `ST_AsMVT()` function uses the function name ("public.countries_name") as the MVT layer name. While this is not required, it allows clients that auto-configure to use the function name as the layer source name.
* In the filter portion of the query (i.e. in the `WHERE` clause), the bounds are transformed to the spatial reference of the table data (in this case, 4326) so that the spatial index on the table geometry can be used.
* In the `ST_AsMVTGeom()` portion of the query, the table geometry is transformed into Web Mercator ([3857](https://epsg.io/3857)) to match the bounds and the _de facto_ expectation that MVT tiles are delivered in Web Mercator projection.
* The `LIMIT` is hard-coded in this example. If you want a user-defined limit, you need to add another parameter to your function definition.
* The function "[volatility](https://www.postgresql.org/docs/current/xfunc-volatility.html)" is declared as `STABLE` because within one transaction context, multiple runs with the same inputs will return the same outputs. It is not marked as `IMMUTABLE` because changes in the base table can change the outputs over time, even for the same inputs.
* The function is declared as `PARALLEL SAFE` because it doesn't depend on any global state that might get confused by running multiple copies of the function at once.
* For earlier versions of PostGIS, the following is an example of a custom function that emulates the behavior of `ST_TileEnvelope()`:
  ```sql
  CREATE OR REPLACE
  FUNCTION ST_TileEnvelope(z integer, x integer, y integer)
  RETURNS geometry
  AS $$
    DECLARE
      size float8;
      zp integer = pow(2, z);
      gx float8;
      gy float8;
    BEGIN
      IF y >= zp OR y < 0 OR x >= zp OR x < 0 THEN
          RAISE EXCEPTION 'invalid tile coordinate (%, %, %)', z, x, y;
      END IF;
      size := 40075016.6855784 / zp;
      gx := (size * x) - (40075016.6855784/2);
      gy := (40075016.6855784/2) - (size * y);
      RETURN ST_SetSRID(ST_MakeEnvelope(gx, gy, gx + size, gy - size), 3857);
    END;
  $$
  LANGUAGE 'plpgsql'
  IMMUTABLE
  STRICT
  PARALLEL SAFE;
  ```

### Spatial processing example

This example clips a layer of [parcels](https://data.vancouver.ca/datacatalogue/propertyInformation.htm) ([EPSG:26910](https://epsg.io/26910)) using a radius and center point, returning only the parcels in the radius, with the boundary parcels clipped to the center.
```sql
CREATE OR REPLACE
FUNCTION public.parcels_in_radius(
                    z integer, x integer, y integer,
                    click_lon float8 default -123.13,
                    click_lat float8 default 49.25,
                    radius float8 default 500.0)
RETURNS bytea
AS $$
    WITH
    args AS (
      SELECT
        ST_TileEnvelope(z, x, y) AS bounds,
        ST_Transform(ST_SetSRID(ST_MakePoint(click_lon, click_lat), 4326), 26910) AS click
    ),
    mvtgeom AS (
      SELECT
        ST_AsMVTGeom(
            ST_Transform(
                ST_Intersection(
                    p.geom,
                    ST_Buffer(args.click, radius)),
                3857),
            args.bounds) AS geom,
        p.site_id
      FROM parcels p, args
      WHERE ST_Intersects(p.geom, ST_Transform(args.bounds, 26910))
      AND ST_DWithin(p.geom, args.click, radius)
      LIMIT 10000
    )
    SELECT ST_AsMVT(mvtgeom, 'public.parcels_in_radius') FROM mvtgeom
$$
LANGUAGE 'sql'
STABLE
PARALLEL SAFE;

COMMENT ON FUNCTION public.parcels_in_radius IS 'Given the click point (click_lon, click_lat) and radius, returns all the parcels in the radius, clipped to the radius circle.';
```
Notes:

* The parcels are stored in a table with spatial reference system [3005](https://epsg.io/3005), a planar projection.
* The click parameters are longitude/latitude, so in building a click geometry (`ST_MakePoint()`) to use for querying, we transform the geometry to the table spatial reference.
* To get the parcel boundaries clipped to the radius, we build a circle in the native spatial reference (26910) using the `ST_Buffer()` function on the click point, then intersect that circle with the parcels.
