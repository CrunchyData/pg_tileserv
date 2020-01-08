# pg_tileserv

[![Travis Build Status][travisbuild]](https://travis-ci.org/CrunchyData/pg_tileserv)

[travisbuild]: https://api.travis-ci.org/CrunchyData/pg_tileserv.svg?branch=master "Travis CI"

An experiment in a [PostGIS](https://postgis.net/)-only tile server in [Go](https://golang.org/). Strip away all the other requirements, it just has to take in HTTP tile requests and form and execute SQL.  In a sincere act of flattery, the API mimics that of the [Martin](https://github.com/urbica/martin) tile server.

# Setup and Installation

## Download

Snapshot builds of the latest code:

* [Linux](https://postgisftw.s3.amazonaws.com/pg_tileserv_snapshot_linux.zip)
* [Windows](https://postgisftw.s3.amazonaws.com/pg_tileserv_snapshot_windows.zip)
* [OSX](https://postgisftw.s3.amazonaws.com/pg_tileserv_snapshot_osx.zip)

## Basic Operation

The simplest start-up uses just a [database connection string](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING) in the `DATABASE_URL` environment variable, and reads all other information from the database.

### Linux/OSX

```sh
export DATABASE_URL=postgresql://username:password@host/dbname
./pg_tileserv
```

### Windows

```
SET DATABASE_URL=postgresql://username:password@host/dbname
pg_tileserv.exe
```

## Configuration File

If you want to alter default values other than the database connection, use the `--config` commandline parameter to pass in a configuration file. In general the defaults are fine, and the program autodetects things like the server name.

```toml
# Database connection
DbConnection = "user=you host=localhost dbname=yourdb"
# Accept connections on this default (default all)
HttpHost = "0.0.0.0"
# Accept connections on this port
HttpPort = 7800
# Advertise URLs relative to this server name
UrlBase = "http://yourserver.com/"
# Resolution to quantize vector tiles to
DefaultResolution = 4096
# Padding to add to vector tiles
DefaultBuffer = 256
# Limit output to this number of features
MaxFeaturesPerTile = 50000
# Advertise this minimum zoom level
DefaultMinZoom = 0
# Advertise this maximum zoom level
DefaultMaxZoom = 22
# Output extra logging information?
Debug = false
```

# Operation

The purpose of `pg_tileserv` is to turn a set of spatial records into tiles, on the fly. The tile server reads two different layers of data:

* Table layers are what they sound like: tables in the database that have a spatial column with a spatial reference system defined on it.
* Function layers hide the source of data from the server, and allow the HTTP client to send in optional parameters to allow more complex SQL functionality. Any function of the form `function(z integer, x integer, y integer, ...)` that returns an MVT `bytea` result can serve as a function layer.

On start-up you can connect to the server and explore the published tables and functions via a web interface at:

* http://localhost:7800

## Layers List

A list of layers is available in JSON at:

* http://localhost:7800/index.json

The index JSON just returns the minimum information about each layer.
```json
{
    "public.ne_50m_admin_0_countries" : {
        "name" : "ne_50m_admin_0_countries",
        "schema" : "public",
        "type" : "table",
        "id" : "public.ne_50m_admin_0_countries",
        "description" : "Natural Earth country data",
        "detailurl" : "http://localhost:7800/public.ne_50m_admin_0_countries.json"
    }
}
```
The `detailurl` provides more detailed metadata for table and function layers.

The `description` field is read from the `comment` value of the table. To set a comment on a table, use the `COMMENT` command.
```sql
COMMENT ON TABLE ne_50m_admin_0_countries IS 'This is my comment';
```

## Table Layers

By default, `pg_tileserv` will provide access to **only** those spatial tables:

* that your database connection has `SELECT` privileges for;
* that include a geometry column;
* that declare a geometry type; and,
* that declare an SRID (spatial reference ID)

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

### Table Detail JSON

In the detail JSON, each layer declares information relevant to setting up a map interface for the layer.
```json
{
   "id" : "public.ne_50m_admin_0_countries",
   "geometrytype" : "MultiPolygon",
   "name" : "ne_50m_admin_0_countries",
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
   "attributes" : [
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
* `id`, `name` and `schema` are the fully qualified, table and schema name of the database table.
* `bounds` and `center` give the extent and middle of the data collection, in geographic coordinates. The order of coordinates in bounds is [minlon, minlat, maxlon, maxlat]. The order of coordinates in center is [lon, lat].
* `tileurl` is the standard substitution pattern URL consumed by map clients like [Mapbox GL JS](https://docs.mapbox.com/mapbox-gl-js/api/) and [OpenLayers](https://openlayers.org).
* `attributes` is a list of attributes in the table, with their data types. The `description` field can be set using the `COMMENT` SQL command:
  ```sql
  COMMENT ON COLUMN ne_50m_admin_0_countries.name_long IS 'This is the long name';
  ```

### Table Tile Request Customization

Most developers will just use the `tileurl` as is, but it possible to add some parameters to the URL to customize behaviour at run time:

* `limit` controls the number of features to write to a tile, the default is 50000.
* `resolution` controls the resolution of a tile, the default is 4096 units per side for a tile.
* `buffer` controls the size of the extra data buffer for a tile, the default is 256 units.
* `attributes` is a comma-separated list of attributes to include in the tile. For wide tables with large numbers of columns, this allows a slimmer tile to be composd.

For example:

    http://localhost:7800/public.ne_50m_admin_0_countries/{z}/{x}/{y}.pbf?limit=100000&attributes=name,long_name

For attribute names that include commas (why did you do that?) [URL encode](https://en.wikipedia.org/wiki/Percent-encoding) the comma in the name string before composing the comma-separated string of all names.

## Function Layers

By default, `pg_tileserv` will provide access to **only** those functions:

* that have `z integer, x integer, y integer` as the first three parameters;
* that return a `bytea`, and
* that your database connection has `EXECUTE` privileges for.

In addition, hopefully obviously, for the function to actually be **useful** it does actually have to return an MVT inside the `bytea` return.

Functions can also have additional parameters to control the generation of tiles: in fact, the whole reason for function layers is to allow **novel dynamic behaviour**.

### Function Layer Examples

#### Filtering Example

This simple example returns just a filtered subset of a table ([ne_50m_admin_0_countries](https://www.naturalearthdata.com/http//www.naturalearthdata.com/download/50m/cultural/ne_50m_admin_0_countries.zip) [EPSG:4326](https://epsg.io/4326)). The filter in this case is the first letters of the name. Note that the `name_prefix` parameter includes a **default value**: this is useful for clients (like the preview interface for this server) that read arbitrary function definitions and need a default value to fill into interface fields.
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
      LIMIT 10000
    )
    SELECT ST_AsMVT(mvtgeom.*, 'public.countries_name') FROM mvtgeom
$$
LANGUAGE 'sql'
STABLE
PARALLEL SAFE;

COMMENT ON FUNCTION public.countries_name IS 'Filters the countries table by the initial letters of the name using the "name_prefix" parameter.';
```
Some notes about this function:

* The `ST_AsMVT()` function uses the function name ("public.countries_name") as the MVT layer name. This is not required, but for clients that self-configure, it allows them to use the function name as the layer source name.
* In the filter portion of the query (in the `WHERE` clause) the bounds are transformed to the spatial reference of the table data (4326) so that the spatial index on the table geometry can be used.
* In the `ST_AsMVTGeom()` portion of the query, the table geometry is transformed into web mercator ([3857](https://epsg.io/3857)) to match the bounds, and the _de facto_ expectation that MVT tiles are delivered in web mercator projection.
* The `ST_TileEnvelope()` function used here is a utility function available in PostGIS 3.0 and higher. For earlier versions, you will probably want to add a custom function to emulate the behavior.
  ```sql
  CREATE OR REPLACE
  FUNCTION TS_TileEnvelope(z integer, x integer, y integer)
  RETURNS geometry
  AS
  $$
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
  STABLE
  PARALLEL SAFE;
  ```
* The `LIMIT` is hard-coded in this example. If you want a user-defined limit you need to add another parameter to your function definition.
* The function "[volatility](https://www.postgresql.org/docs/current/xfunc-volatility.html)" is declared as `STABLE` because within one transaction context, multiple runs with the same inputs will return the same outputs. It is not marked as `IMMUTABLE` because changes in the base table can change the outputs over time, even for the same inputs.
* The function is declared as `PARALLEL SAFE` because it doesn't depend on any global state that might get confused by running multiple copies of the function at once.

#### Spatial Processing Example

This example clips a layer of [parcels](https://data.vancouver.ca/datacatalogue/propertyInformation.htm) [EPSG:26910](https://epsg.io/26910) using a radius and center point, returning only the parcels in the radius, with the boundary parcels clipped to the center.
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
    SELECT ST_AsMVT(mvtgeom.*, 'public.parcels_in_radius') FROM mvtgeom
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

#### Dynamic Geometry Example

So far all our examples have used simple SQL functions, but using the more [procedural PL/PgSQL language](https://www.postgresql.org/docs/current/plpgsql.html) we can create much more interactive examples.

```sql
CREATE OR REPLACE
FUNCTION public.squares(z integer, x integer, y integer, depth integer default 2)
RETURNS bytea
AS $$
DECLARE
    result bytea;
    sq_width float8;
    tile_xmin float8;
    tile_ymin float8;
    bounds geometry;
BEGIN
    -- Find the tile bounds
    SELECT ST_TileEnvelope(z, x, y) AS geom INTO bounds;
    -- Find the bottom corner of the bounds
    tile_xmin := ST_XMin(bounds);
    tile_ymin := ST_YMin(bounds);
    -- We want tile divided up into depth*depth squares per tile,
    -- so what is the width of a square?
    sq_width := (ST_XMax(bounds) - ST_XMin(bounds)) / depth;

    WITH mvtgeom AS (
        SELECT
            -- Fill in the tile with all the squares
            ST_AsMVTGeom(ST_MakeEnvelope(
                tile_xmin + sq_width * (a-1),
                tile_ymin + sq_width * (b-1),
                tile_xmin + sq_width * a,
                tile_ymin + sq_width * b), bounds),
            -- Each square gets an attribute that shows
            -- what tile it is a part of and what its sub-address
            -- in that tile is
            Format('(%s.%s,%s.%s)', x, a, y, b) AS tilecoord
        -- Drive the square generator with a two-dimensional
        -- generate_series setup
        FROM generate_series(1, depth) a, generate_series(1, depth) b
        )
    SELECT ST_AsMVT(mvtgeom.*, 'public.squares')
    -- Put the query result into the result variale.
    INTO result FROM mvtgeom;

    -- Return the answer
    RETURN result;
END;
$$
LANGUAGE 'plpgsql'
IMMUTABLE -- Same inputs always give same outputs
STRICT -- Null input gets null output
PARALLEL SAFE;

COMMENT ON FUNCTION public.squares IS 'For each tile requested, generate and return depth*depth polygons covering the tile. The effect is one of always having a grid coverage at the appropriate current scale.';
```

#### Dynamic Geometry with Spatial Join Example

**TO BE DONE**

```sql
CREATE OR REPLACE
FUNCTION Hexagon(i integer, j integer, edge float8)
RETURNS geometry
AS $$
    WITH t AS (SELECT edge AS e, edge*cos(pi()/6) AS h)
    SELECT
        ST_MakePolygon(ST_MakeLine(ARRAY[
            ST_MakePoint(1.5*i*e - 1.0*e, h*(2*j+(i%2)) + 0),
            ST_MakePoint(1.5*i*e - 0.5*e, h*(2*j+(i%2)) + -1*h),
            ST_MakePoint(1.5*i*e + 0.5*e, h*(2*j+(i%2)) + -1*h),
            ST_MakePoint(1.5*i*e + 1.0*e, h*(2*j+(i%2)) + 0),
            ST_MakePoint(1.5*i*e + 0.5*e, h*(2*j+(i%2)) + h),
            ST_MakePoint(1.5*i*e - 0.5*e, h*(2*j+(i%2)) + h),
            ST_MakePoint(1.5*i*e - 1.0*e, h*(2*j+(i%2)) + 0)
        ]))
    FROM t
$$
LANGUAGE 'sql'
IMMUTABLE
STRICT
PARALLEL SAFE;

CREATE OR REPLACE
FUNCTION HexagonCoordinates(bounds geometry, edge float8, OUT i integer, OUT j integer)
RETURNS SETOF record
AS $$
    DECLARE
        mini integer;
        maxi integer;
        minj integer;
        maxj integer;
        h float8 := edge*cos(pi()/6);
    BEGIN
    mini := floor(st_xmin(bounds) / (1.5*edge));
    minj := floor(st_ymin(bounds) / (2*h));
    maxi := ceil(st_xmax(bounds) / (1.5*edge));
    maxj := ceil(st_ymax(bounds) / (2*h));
    FOR i, j IN
    SELECT a, b
    FROM generate_series(mini, maxi) a,
         generate_series(minj, maxj) b
    LOOP
        RETURN NEXT;
    END LOOP;
    END;
$$
LANGUAGE 'plpgsql'
IMMUTABLE
STRICT
PARALLEL SAFE;

CREATE OR REPLACE
FUNCTION TileHexagons(z integer, x integer, y integer, step integer,
                      OUT geom geometry(Polygon, 3857), OUT i integer, OUT j integer)
RETURNS SETOF record
AS $$
    DECLARE
        bounds geometry;
        maxbounds geometry := ST_TileEnvelope(0, 0, 0);
        edge float8;
    BEGIN
    bounds := ST_TileEnvelope(z, x, y);
    edge := (ST_XMax(bounds) - ST_XMin(bounds)) / pow(2, step);
    FOR geom, i, j IN
    SELECT ST_SetSRID(Hexagon(h.i, h.j, edge), 3857), h.i, h.j
    FROM HexagonCoordinates(bounds, edge) h
    LOOP
        IF maxbounds ~ geom AND bounds && geom THEN
            RETURN NEXT;
        END IF;
    END LOOP;
    END;
$$
LANGUAGE 'plpgsql'
IMMUTABLE
STRICT
PARALLEL SAFE;
```

```sql
CREATE OR REPLACE
FUNCTION HexPopulationSummary(z integer, x integer, y integer, step integer default 4)
RETURNS bytea
AS $$
WITH
bounds AS (
    SELECT ST_TileEnvelope(z, x, y) AS geom
),
rows AS (
    SELECT Sum(pop_max) AS pop_max, Sum(pop_min) AS pop_min, h.i, h.j, h.geom
    FROM TileHexagons(z, x, y, step) h
    JOIN ne_50m_populated_places n
    ON ST_Intersects(n.geom, ST_Transform(h.geom, 4326))
    GROUP BY h.i, h.j, h.geom
),
mvt AS (
    SELECT ST_AsMVTGeom(rows.geom, bounds.geom) AS geom,
           rows.pop_max, rows.pop_min, rows.i, rows.j
    FROM rows, bounds
)
SELECT ST_AsMVT(mvt.*, 'hexes') FROM mvt
$$
LANGUAGE 'sql';
```

```sql
CREATE OR REPLACE
FUNCTION HexPopulationSummary3(z integer, x integer, y integer, arg1 text default 'arg1', arg2 integer default 101)
RETURNS bytea
AS $$
WITH
bounds AS (
    SELECT ST_TileEnvelope(z, x, y) AS geom
),
rows AS (
    SELECT Sum(pop_max) AS pop_max, Sum(pop_min) AS pop_min, h.i, h.j, h.geom
    FROM TileHexagons(z, x, y, 4) h
    JOIN ne_50m_populated_places n
    ON ST_Intersects(n.geom, ST_Transform(h.geom, 4326))
    GROUP BY h.i, h.j, h.geom
),
mvt AS (
    SELECT ST_AsMVTGeom(rows.geom, bounds.geom) AS geom,
           rows.pop_max, rows.pop_min, rows.i, rows.j
    FROM rows, bounds
)
SELECT ST_AsMVT(mvt.*, 'hexes') FROM mvt
$$
LANGUAGE 'sql';
```

```sql
CREATE FUNCTION foobar(integer, b integer default 4, c text default 'ghgh', e geometry default 'Point(0 0)'::geometry(point, 4326)) returns integer as 'select $1 + $2' language 'sql';


SELECT
Format('%s.%s', n.nspname, p.proname) AS id,
n.nspname,
p.proname,
d.description,
p.proargnames AS argnames,
string_to_array(oidvectortypes(p.proargtypes),', ') AS argtypes
FROM pg_proc p
JOIN pg_namespace n ON (p.pronamespace = n.oid)
LEFT JOIN pg_description d ON (p.oid = d.objoid)
WHERE p.proargtypes[0:2] = ARRAY[23::oid, 23::oid, 23::oid]
AND p.proargnames[1:3] = ARRAY['z'::text, 'x'::text, 'y'::text]
AND prorettype = 17
AND has_function_privilege(Format('%s.%s(%s)', n.nspname, p.proname, oidvectortypes(proargtypes)), 'execute') ;
```

# Testing

* table tile
  * limit specified
  * one attribute specified
  * all attributes specified
  * non-existing attribute specified
* geometry only table
* geometry and pk only table
