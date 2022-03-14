<p align="center">
  <a href="https://access.crunchydata.com/documentation/pg_tileserv/latest/"><img width="180" height="180" src="./hugo/static/crunchy-spatial-logo.png?raw=true" /></a>
</p>

# pg_tileserv

[![.github/workflows/ci.yml](https://github.com/CrunchyData/pg_tileserv/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/CrunchyData/pg_tileserv/actions/workflows/ci.yml)

A [PostGIS](https://postgis.net/)-only tile server in [Go](https://golang.org/). Strip away all the other requirements, it just has to take in HTTP tile requests and form and execute SQL.  In a sincere act of flattery, the API looks a lot like that of the [Martin](https://github.com/urbica/martin) tile server.

* https://access.crunchydata.com/documentation/pg_tileserv/latest/

# Setup and Installation

## Download

Builds of the latest code:

* [Linux](https://postgisftw.s3.amazonaws.com/pg_tileserv_latest_linux.zip)
* [Windows](https://postgisftw.s3.amazonaws.com/pg_tileserv_latest_windows.zip)
* [MacOS](https://postgisftw.s3.amazonaws.com/pg_tileserv_latest_macos.zip)
* [Docker](https://hub.docker.com/r/pramsey/pg_tileserv)

## Basic Operation

The executable will read user/connection information from the `DATABASE_URL` and connect to the database, exposing all functions and tables the database user has read and execute permissions on.

For **production deployment**, place an HTTP proxy caching layer (eg [Varnish](https://varnish-cache.org/)) in between the tile server and clients to reduce database load and increase application performance.

### Linux/MacOS

```sh
export DATABASE_URL=postgresql://username:password@host/dbname
./pg_tileserv
```

### Windows

```
SET DATABASE_URL=postgresql://username:password@host/dbname
pg_tileserv.exe
```

### Docker

Use [Dockerfile.alpine](Dockerfile.alpine) to build a lightweight (18MB expanded) Docker Image.
See also [a full example with Docker Compose](examples/docker/README.md).

## Trouble-shooting

To get more information about what is going on behind the scenes, run with the `--debug` commandline parameter on, or turn on debugging in the configuration file:
```sh
./pg_tileserv --debug
```

## Configuration File

The configuration file will be automatically read from the following locations, if it exists:

* In the system configuration directory, at `/etc/pg_tileserv.toml`
* Relative to the directory from which the program is run, `./config/pg_tileserv.toml`
* In a root volume at `/config/pg_tileserv.toml`

If you want to pass a path directly to the configuration file, use the `--config` commandline parameter to pass in a pull path to configuration file. When using the `--config` option, configuration files in other locations will be ignored.

```sh
./pg_tileserv --config /opt/pg_tileserv/pg_tileserv.toml
```

In general the defaults are fine, and the program autodetects things like the server name.

```toml
# Database connection
DbConnection = "user=you host=localhost dbname=yourdb"

# Close pooled connections after this interval
DbPoolMaxConnLifeTime = "1h"

# Hold no more than this number of connections in the database pool
DbPoolMaxConns = 4

# Look to read html templates from this directory
AssetsPath = "/usr/share/pg_tileserv/assets"

# Accept connections on this subnet (default accepts on all)
HttpHost = "0.0.0.0"

# Accept connections on this port
HttpPort = 7800
HttpsPort = 7801

# HTTPS configuration
# TLS server certificate full chain and private key
# If you do not specify both, the TLS server will not be started
TlsServerCertificateFile = "server.crt"
TlsServerPrivateKeyFile = "server.key"
```

For SSL support, you will need both a server private key and an authority certificate. For testing purposes you can generate a self-signed key/cert pair using `openssl`:

```bash
openssl req  -nodes -new -x509  -keyout server.key -out server.crt
```

```toml
# Cache control configuration. TTL is time in seconds to request
# that responses be cached by any downstream caching services.
# Zero means no cache control header will be set.
CacheTTL = 60

# Advertise URLs relative to this server name
# default is to looke this up from incoming request headers
# UrlBase = "http://yourserver.com/"
# Resolution to quantize vector tiles to
DefaultResolution = 4096
# Padding to add to vector tiles
DefaultBuffer = 256
# Limit number of features requested (-1 = no limit)
MaxFeaturesPerTile = 10000
# Advertise this minimum zoom level
DefaultMinZoom = 0
# Advertise this maximum zoom level
DefaultMaxZoom = 22

# Allow any page to consume these tiles
CORSOrigins = ["*"]

# Output extra logging information?
Debug = false

# Enable Prometheus metrics
# Metrics will be exported at `/metrics`.
EnableMetrics = false

# Default CS is Web Mercator (EPSG:3857)
[CoordinateSystem]
SRID = 3857
Xmin = -20037508.3427892
Ymin = -20037508.3427892
Xmax = 20037508.3427892
Ymax = 20037508.3427892
```
You can use the **CoordinateSystem** block to output files in a system other than the default [Web Mercator](http://epsg.io/3857) projection. In order to view a map with multiple layers in a non-standard projection, you will have to ensure that all layers share the same projection, otherwise the layers will not line up.

### Configuration Using Environment Variables

Any parameter in the configuration file can be over-ridden at run-time in the environment. Prepend the upper-cased parameter name with `TS_` to set the value. For example, to change the HTTP port using the environment:
```bash
export TS_HTTPPORT=8889
```


# Operation

The purpose of `pg_tileserv` is to turn a set of spatial records into tiles, on the fly. The tile server reads two different layers of data:

* Table layers are what they sound like: tables in the database that have a spatial column with a spatial reference system defined on it.
* Function layers hide the source of data from the server, and allow the HTTP client to send in optional parameters to allow more complex SQL functionality. Any function of the form `function(z integer, x integer, y integer, ...)` that returns an MVT `bytea` result can serve as a function layer.

## Web Interface

After start-up you can connect to the server and explore the published tables and functions in the database via a web interface at:

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
* Connect `pg_tileserv` to the database using that role


If your table contains a geometry column that appears valid, but it is not
available within `pg_tileserv`, you may need to specifically set a geometry
type or SRID.

To determine if a table is compatible, make sure that it is returned by the
following query:

```sql
SELECT
	nspname AS SCHEMA,
	relname AS TABLE,
	attname AS geometry_column,
	postgis_typmod_srid (atttypmod) AS srid,
	postgis_typmod_type (atttypmod) AS geometry_type
FROM
	pg_class c
	JOIN pg_namespace n ON (c.relnamespace = n.oid)
	JOIN pg_attribute a ON (a.attrelid = c.oid)
	JOIN pg_type t ON (a.atttypid = t.oid)
WHERE
	relkind IN('r', 'v', 'm')
	AND typname = 'geometry'
    AND postgis_typmod_srid (atttypmod) != 0
	AND relname = '<mytable>';
```

If not, make sure that the geometry column has a valid SRID defined in the table
metadata. You may need to specifically assign a geometry type, especially if the
table was created using a `SELECT` query from another geometry table.

For example, to set the geometry as a `Point` type:

```SQL
ALTER TABLE mytable ALTER COLUMN geom TYPE geometry (Point, 4326);
```

### Table Layer Detail JSON

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
* `id`, `name` and `schema` are the fully qualified, table and schema name of the database table.
* `bounds` and `center` give the extent and middle of the data collection, in geographic coordinates. The order of coordinates in bounds is [minlon, minlat, maxlon, maxlat]. The order of coordinates in center is [lon, lat].
* `tileurl` is the standard substitution pattern URL consumed by map clients like [Mapbox GL JS](https://docs.mapbox.com/mapbox-gl-js/api/) and [OpenLayers](https://openlayers.org).
* `properties` is a list of columns in the table, with their data types and descriptions. The column `description` field can be set using the `COMMENT` SQL command, for example:
  ```sql
  COMMENT ON COLUMN ne_50m_admin_0_countries.name_long IS 'This is the long name';
  ```

### Feature id
The [vector tile specification](https://github.com/mapbox/vector-tile-spec/tree/master/2.1#42-features) allows for an *optional* `id` field for each feature. This field should be unique within the parent layer.

By default, `pg_tileserv` will generate this id if:

* the PostGIS version is >= 3.0 and;
* the table has a primary key and;
* the primary key field is one of ``'int2', 'int4', 'int8'``

A feature id will not be generated for Views since these do not have a primary key. In cases where an `id` is not generated, depending on the map renderer, it may be possible to generate a feature id at runtime. See https://docs.mapbox.com/mapbox-gl-js/style-spec/sources/#vector-promoteId for an example in Mapbox GL JS.

### Table Tile Request Customization

Most developers will just use the `tileurl` as is, but it possible to add some parameters to the URL to customize behaviour at run time:

* `limit` controls the number of features to write to a tile, the default is 50000.
* `resolution` controls the resolution of a tile, the default is 4096 units per side for a tile.
* `buffer` controls the size of the extra data buffer for a tile, the default is 256 units.
* `properties` is a comma-separated list of properties to include in the tile. For wide tables with large numbers of columns, this allows a slimmer tile to be composed.
* `filter` is a CQL logical expression which specifies the features to be included in the tile.  See the [CQL documentation](hugo/content/usage/cql.md).

For example:

    http://localhost:7800/public.ne_50m_admin_0_countries/{z}/{x}/{y}.pbf?limit=100000&properties=name,long_name

For property names that include commas (why did you do that?) [URL encode](https://en.wikipedia.org/wiki/Percent-encoding) the comma in the name string before composing the comma-separated string of all names.

### Multi-Layer Tile Requests

For more complex applications, multi-layer tiles can be useful to cut down on the amount of HTTP requests to pull in vector tiles. Doing this with `pg_tileserv` is easy, just add additional tables to your request. You can add as many tables as you like to your request, just separate them with a comma.

For example:

    http://localhost:7800/public.ne_50m_admin_0_countries,public.ne_50m_airports/{z}/{x}/{y}.pbf

## Function Layers

By default, `pg_tileserv` will provide access to **only** those functions:

* that have `z integer, x integer, y integer` as the first three parameters;
* that return a `bytea`, and
* that your database connection has `EXECUTE` privileges for.

In addition, hopefully obviously, for the function to actually be **useful** it does actually have to return an MVT inside the `bytea` return.

Functions can also have additional parameters to control the generation of tiles: in fact, the whole reason for function layers is to allow **novel dynamic behaviour**.

### Function Layer Detail JSON

In the detail JSON, each function declares information relevant to setting up a map interface for the layer. Because functions generate tiles dynamically, the system cannot auto-discover things like extent or center, unfortunately. However, the custom parameters and defaults can be read from the function definition and exposed in the detail JSON.
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
* `description` can be set using `COMMENT ON FUNCTION` SQL command.
* `id`, `schema` and `name` are the fully qualified name, schema and function name, respectively.
* `minzoom` and `maxzoom` are just the defaults, as set in the configuration file.
* `arguments` is a list of argument names, with the data type and default value.

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
DECLARE
    result bytea;
BEGIN
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
    SELECT ST_AsMVT(mvtgeom, 'public.countries_name')
    INTO result
    FROM mvtgeom;

    RETURN result;
END;
$$
LANGUAGE 'plpgsql'
STABLE
PARALLEL SAFE;

COMMENT ON FUNCTION public.countries_name IS 'Filters the countries table by the initial letters of the name using the "name_prefix" parameter.';
```
Some notes about this function:

* The `ST_AsMVT()` function uses the function name ("public.countries_name") as the MVT layer name. This is not required, but for clients that self-configure, it allows them to use the function name as the layer source name.
* In the filter portion of the query (in the `WHERE` clause) the bounds are transformed to the spatial reference of the table data (4326) so that the spatial index on the table geometry can be used.
* In the `ST_AsMVTGeom()` portion of the query, the table geometry is transformed into web mercator ([3857](https://epsg.io/3857)) to match the bounds, and the _de facto_ expectation that MVT tiles are delivered in web mercator projection.
* The `LIMIT` is hard-coded in this example. If you want a user-defined limit you need to add another parameter to your function definition.
* The function "[volatility](https://www.postgresql.org/docs/current/xfunc-volatility.html)" is declared as `STABLE` because within one transaction context, multiple runs with the same inputs will return the same outputs. It is not marked as `IMMUTABLE` because changes in the base table can change the outputs over time, even for the same inputs.
* The function is declared as `PARALLEL SAFE` because it doesn't depend on any global state that might get confused by running multiple copies of the function at once.
* The `ST_TileEnvelope()` function used here is a utility function available in PostGIS 3.0 and higher. For earlier versions, you will probably want to add a custom function to emulate the behavior.
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
DECLARE
    result bytea;
BEGIN
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
    SELECT ST_AsMVT(mvtgeom, 'public.parcels_in_radius')
    INTO result
    FROM mvtgeom;

    RETURN result;
END;
$$
LANGUAGE 'plpgsql'
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
            -- Each square gets a property that shows
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

#### Dynamic Hexagons with Spatial Join Example

Hexagonal tilings are popular with data visualization experts because they can be used to summarize point data without adding a visual bias to the output via different summary area sizes. They also have a nice "non-pointy" shape, while still providing a complete tiling of the plane.

When you want to provide a hexagonal summary of a data set at multiple scales, you have an implementation problem: do you need to create a pile of hexagon tables, solely for the purpose of summary visualization?

No, you don't have to, you can generate your hexagons dynamically based on the scale of the requested map tiles.

The first challenge is that a hexagon tile set cannot be perfectly inscribed into a powers-of-two square tile set. That means that any given tile will contain some odd combination of full and partial hexagons. In order for the hexagons that straddle tile boundaries to match up, we need a hexagon tiling that is uniform over the whole plane.

So, our first function takes a "hexagon grid coordinate" and generates a hexagon for that coordinate. The size and location of that hexagon are controlled by the hexagon edge length for this particular tiling.
```sql
-- Given coordinates in the hexagon tiling that has this
-- edge size, return the built-out hexagon
CREATE OR REPLACE
FUNCTION hexagon(i integer, j integer, edge float8)
RETURNS geometry
AS $$
DECLARE
h float8 := edge*cos(pi()/6.0);
cx float8 := 1.5*i*edge;
cy float8 := h*(2*j+abs(i%2));
BEGIN
RETURN ST_MakePolygon(ST_MakeLine(ARRAY[
            ST_MakePoint(cx - 1.0*edge, cy + 0),
            ST_MakePoint(cx - 0.5*edge, cy + -1*h),
            ST_MakePoint(cx + 0.5*edge, cy + -1*h),
            ST_MakePoint(cx + 1.0*edge, cy + 0),
            ST_MakePoint(cx + 0.5*edge, cy + h),
            ST_MakePoint(cx - 0.5*edge, cy + h),
            ST_MakePoint(cx - 1.0*edge, cy + 0)
        ]));
END;
$$
LANGUAGE 'plpgsql'
IMMUTABLE
STRICT
PARALLEL SAFE;

SELECT ST_AsText(hexagon(2, 2, 10.0));
```
```
 POLYGON((20 34.6410161513775,25 25.9807621135332,
          35 25.9807621135332,40 34.6410161513775,
          35 43.3012701892219,25 43.3012701892219,
          20 34.6410161513775))
```
Now we need a function that, given a square input (a map tile) can figure out all the hexagon coordinates that fall within the tile. Again, the edge size of the hexagon tiling determines the overall geometry of the hex tiling. More than one hexagon will be required, most times, so this is a set-returning function.
```sql
-- Given a square bounds, find all the hexagonal cells
-- of a hex tiling (determined by edge size)
-- that might cover that square (slightly over-determined)
CREATE OR REPLACE
FUNCTION hexagoncoordinates(bounds geometry, edge float8,
                            OUT i integer, OUT j integer)
RETURNS SETOF record
AS $$
    DECLARE
        h float8 := edge*cos(pi()/6);
        mini integer := floor(st_xmin(bounds) / (1.5*edge));
        minj integer := floor(st_ymin(bounds) / (2*h));
        maxi integer := ceil(st_xmax(bounds) / (1.5*edge));
        maxj integer := ceil(st_ymax(bounds) / (2*h));
    BEGIN
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

SELECT * FROM hexagoncoordinates(ST_TileEnvelope(15, 1, 1), 1000.0);
```
```
   i    |   j
--------+-------
 -13358 | 11567
 -13358 | 11568
 -13357 | 11567
 -13357 | 11568
 -13356 | 11567
 -13356 | 11568
```
Next, a function that puts the two parts together. With tile coordinates and edge size as input, generate the set of all the hexagons that cover the tile. The output here is basically a spatial table: a set of rows, each row containing a geometry (hexagon) and some properties (hexagon coordinates). Just the input we need for a spatial join.
```sql
-- Given an input ZXY tile coordinate, output a set of hexagons
-- (and hexagon coordinates) in web mercator that cover that tile
CREATE OR REPLACE
FUNCTION tilehexagons(z integer, x integer, y integer, step integer,
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
    SELECT ST_SetSRID(hexagon(h.i, h.j, edge), 3857), h.i, h.j
    FROM hexagoncoordinates(bounds, edge) h
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
The function that the tile server actually calls looks like all other tile server functions: tile coordinates and optional parameter input; `bytea` MVT output.
```sql
-- Given an input tile, generate the covering hexagons,
-- spatially join to population table, summarize
-- population in each hexagon, and generate MVT
-- output of the result. Step parameter determines
-- how many hexagons to generate per tile.
CREATE OR REPLACE
FUNCTION public.hexpopulationsummary(z integer, x integer, y integer, step integer default 4)
RETURNS bytea
AS $$
DECLARE
    result bytea;
BEGIN
    WITH
    bounds AS (
        -- Convert tile coordinates to web mercator tile bounds
        SELECT ST_TileEnvelope(z, x, y) AS geom
    ),
    rows AS (
        -- Summary of populated places grouped by hex
        SELECT Sum(pop_max) AS pop_max, Sum(pop_min) AS pop_min, h.i, h.j, h.geom
        -- All the hexes that interact with this tile
        FROM TileHexagons(z, x, y, step) h
        -- All the populated places
        JOIN ne_50m_populated_places n
        -- Transform the hex into the SRS (4326 in this case)
        -- of the table of interest
        ON ST_Intersects(n.geom, ST_Transform(h.geom, 4326))
        GROUP BY h.i, h.j, h.geom
    ),
    mvt AS (
        -- Usual tile processing, ST_AsMVTGeom simplifies, quantizes,
        -- and clips to tile boundary
        SELECT ST_AsMVTGeom(rows.geom, bounds.geom) AS geom,
               rows.pop_max, rows.pop_min, rows.i, rows.j
        FROM rows, bounds
    )
    -- Generate MVT encoding of final input record
    SELECT ST_AsMVT(mvt, 'public.hexpopulationsummary')
    INTO result
    FROM mvt;

    RETURN result;
END:
$$
LANGUAGE 'plpgsql'
STABLE
STRICT
PARALLEL SAFE;

COMMENT ON FUNCTION public.hexpopulationsummary IS 'Hex summary of the ne_50m_populated_places table. Step parameter determines how approximately many hexes (2^step) to generate per tile.';
```
A basic "just hexes" layer that skips the spatial join step is even simpler.
```sql
-- Given an input tile, generate the covering hexagons Step parameter determines
-- how many hexagons to generate per tile.
CREATE OR REPLACE
FUNCTION public.hexagons(z integer, x integer, y integer, step integer default 4)
RETURNS bytea
AS $$
DECLARE
    result bytea;
BEGIN
    WITH
    bounds AS (
        -- Convert tile coordinates to web mercator tile bounds
        SELECT ST_TileEnvelope(z, x, y) AS geom
    ),
    rows AS (
        -- All the hexes that interact with this tile
        SELECT h.i, h.j, h.geom
        FROM TileHexagons(z, x, y, step) h
    ),
    mvt AS (
        -- Usual tile processing, ST_AsMVTGeom simplifies, quantizes,
        -- and clips to tile boundary
        SELECT ST_AsMVTGeom(rows.geom, bounds.geom) AS geom,
               rows.i, rows.j
        FROM rows, bounds
    )
    -- Generate MVT encoding of final input record
    SELECT ST_AsMVT(mvt, 'public.hexagons')
    INTO result
    FROM mvt;

    RETURN result;
END;
$$
LANGUAGE 'plpgsql'
STABLE
STRICT
PARALLEL SAFE;

COMMENT ON FUNCTION public.hexagons IS 'Hex coverage dynamically generated. Step parameter determines how approximately many hexes (2^step) to generate per tile.';
```

# Security

The basic principle of security is to connect your tile server to the database with a user that has just the access you want it to have, and no more. To support different access patterns, create different users with access to different tables/functions, and run multiple services, connecting with those different users.
```sql
CREATE USER tileserver;
```
Start with a blank user. A blank user will have no select privileges on tables it does not own. It will have execute privileges on functions. However, any the user will have no select privileges on tables accessed by functions, so effectively the user will still have no access to data.

## Tables

If your tables are in a schema other than public, you will have to also grant "usage" on that schema to your user.
```sql
GRANT USAGE ON SCHEMA myschema TO tileserver;
```
You can then grant access to the user one table at a time.
```sql
GRANT SELECT ON TABLE myschema.mytable TO tileserver;
```
Or grant access to all the tables at once.
```sql
GRANT SELECT ON ALL TABLES IN SCHEMA myschema TO tileserver;
```

## Functions

As noted above, functions that access table data effectively are restricted by the access levels the user has to the tables the function reads. However, if you want to completely restrict access to the function, including visibility in the user interface, you can strip execution privileges from the function.

```sql
-- All functions grant execute to 'public' and all roles are
-- part of the 'public' group, so public has to be removed
-- from the executors of the function
REVOKE EXECUTE ON FUNCTION myschema.myfunction FROM public;
-- Just to be sure, also revoke execute from the user
REVOKE EXECUTE ON FUNCTION myschema.myfunction FROM tileserver;
```
