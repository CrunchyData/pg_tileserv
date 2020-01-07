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

The simplest start-up uses just a [database connection string](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING) in an environment variable, and reads all other information from the database.

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

  http://localhost:7800

## Layers List

A list of layers is available in JSON at:

  http://localhost:7800/index.json

Each layer declares information relevant to setting up a map layer for the source: a web client should be able to self-configure using the information in the layer 

## Table Layers

By default, `pg_tileserv` will provide access to **only** those spatial tables:

* that your database connection has `SELECT` privileges for;
* that include a geometry column
* that declare a geometry type
* that declare an SRID (spatial reference ID)

To restrict access to a certain set of tables, use database security principles:

* Create a role with limited privileges
* Only grant `SELECT` to that role for tables you want to publish
* Only grant `EXECUTE` to that role for functions you want to publish

## Function Layers

**FROM HERE DOWN IS TO BE DONE**

```sql
CREATE OR REPLACE
FUNCTION lakes(z integer, x integer, y integer, name_prefix text default '')
RETURNS bytea
AS $$
    WITH
    bounds AS (
      SELECT ST_TileEnvelope(z, x, y) AS geom
    ),
    mvtgeom AS (
      SELECT ST_AsMVTGeom(ST_Transform(t.geom, 3857), bounds.geom) AS geom,
        t.name
      FROM ne_50m_lakes t, bounds
      WHERE ST_Intersects(t.geom, ST_Transform(bounds.geom, 4326))
      AND t.name like (name_prefix || '%')
      LIMIT 10000
    )
    SELECT ST_AsMVT(mvtgeom.*, 'public.lakes') FROM mvtgeom
$$
LANGUAGE 'sql';
```


```sql
CREATE OR REPLACE
FUNCTION public.squares(z integer, x integer, y integer, depth integer default 2)
RETURNS bytea
AS $$
DECLARE
rslt bytea;
width float8;
xmin float8;
ymin float8;
bounds geometry;
BEGIN
    -- Get tile bounds
    SELECT ST_TileEnvelope(z, x, y) AS geom INTO bounds;
    xmin := ST_XMin(bounds);
    ymin := ST_YMin(bounds);
    width := (ST_XMax(bounds) - ST_XMin(bounds)) / depth;
    WITH mvtgeom AS (
        SELECT ST_AsMVTGeom(ST_ExteriorRing(ST_MakeEnvelope(
            xmin + width * (a-1), ymin + width * (b-1),
            xmin + width * a, ymin + width * b)), bounds),
            Format('(%s.%s,%s.%s)', x, a, y, b) AS tilecoord
        FROM generate_series(1, depth) a, generate_series(1, depth) b
        )
    SELECT ST_AsMVT(mvtgeom.*, 'public.squares')
    INTO rslt FROM mvtgeom;
    RETURN rslt;
END;
$$
LANGUAGE 'plpgsql'
IMMUTABLE
STRICT
PARALLEL SAFE;
```







```
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

```
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

```
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
