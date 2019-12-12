# pg_tileserv

An experiment in a [PostGIS](https://postgis.net/)-only tile server in [Go](https://golang.org/). Strip away all the other requirements, it just has to take in HTTP tile requests and form and execute SQL.  In a sincere act of flattery, I have mostly copied the API of the [Martin](https://github.com/urbica/martin) tile server.

## Table Sources



## Function Sources

```
CREATE OR REPLACE
FUNCTION countries_name(z integer, x integer, y integer, name_prefix text)
RETURNS bytea
AS $$
    WITH
    bounds AS (
      SELECT ST_TileEnvelope(z, x, y) AS geom
    ),
    mvtgeom AS (
      SELECT ST_AsMVTGeom(ST_Transform(t.geom, 3857), bounds.geom) AS geom, t.name
      FROM ne_50m_admin_0_countries t, bounds
      WHERE ST_Intersects(t.geom, ST_Transform(bounds.geom, 4326))
      AND t.name like (name_prefix || '%')
      LIMIT 10000
    )
    SELECT ST_AsMVT(mvtgeom.*, 'ne_50m_admin_0_countries') FROM mvtgeom
$$
LANGUAGE 'sql';
```

```
CREATE OR REPLACE
FUNCTION squares(z integer, x integer, y integer, depth integer)
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
        SELECT ST_AsMVTGeom(ST_MakeEnvelope(
            xmin + width * (a-1), ymin + width * (b-1),
            xmin + width * a, ymin + width * b), bounds),
            Format('(%s.%s,%s.%s)', x, a, y, b) AS tilecoord
        FROM generate_series(1, depth) a, generate_series(1, depth) b
        )
    SELECT ST_AsMVT(mvtgeom.*, 'tile_grid')
    INTO rslt FROM mvtgeom;
    RETURN rslt;
END;
$$
LANGUAGE 'plpgsql';
```


```
CREATE OR REPLACE
FUNCTION hexsummary(z integer, x integer, y integer, table_name text, depth integer default 4)
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
    width := (ST_XMax(bounds) - ST_XMin(bounds))
     / depth;
    WITH mvtgeom AS (
        SELECT ST_AsMVTGeom(ST_MakeEnvelope(
            xmin + width * (a-1), ymin + width * (b-1),
            xmin + width * a, ymin + width * b), bounds),
            Format('(%s.%s,%s.%s)', x, a, y, b) AS tilecoord
        FROM generate_series(1, depth) a, generate_series(1, depth) b
        )
    SELECT ST_AsMVT(mvtgeom.*, 'tile_grid')
    INTO rslt FROM mvtgeom;
    RETURN rslt;
END;
$$
LANGUAGE 'plpgsql';
```


CREATE OR REPLACE
FUNCTION Hexagon(i integer, j integer, edge float8)
RETURNS geometry
AS $$
WITH t AS (SELECT edge, edge*cos(pi()/6) AS h)
SELECT
ST_MakePolygon(ST_MakeLine(
ARRAY[
ST_MakePoint(1.5*i*e - 1.0*e, h*(2*j+(i%2)) + 0),
ST_MakePoint(1.5*i*e - 0.5*e, h*(2*j+(i%2)) + -1*h),
ST_MakePoint(1.5*i*e + 0.5*e, h*(2*j+(i%2)) + -1*h),
ST_MakePoint(1.5*i*e + 1.0*e, h*(2*j+(i%2)) + 0),
ST_MakePoint(1.5*i*e + 0.5*e, h*(2*j+(i%2)) + h),
ST_MakePoint(1.5*i*e - 0.5*e, h*(2*j+(i%2)) + h),
ST_MakePoint(1.5*i*e - 1.0*e, h*(2*j+(i%2)) + 0)
]
))
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
FROM generate_series(mini, maxi) a, generate_series(minj, maxj) b
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
FUNCTION TileHexagons(z integer, x integer, y integer, step integer)
RETURNS SETOF geometry
AS $$
DECLARE
bounds geometry;
maxbounds geometry := ST_TileEnvelope(0, 0, 0);
edge float8;
g geometry;
BEGIN
bounds := ST_TileEnvelope(z, x, y);
edge := (ST_XMax(bounds) - ST_XMin(bounds)) / pow(2, step);
FOR g IN
SELECT ST_SetSRID(Hexagon(i, j, edge), 3857)
FROM HexagonCoordinates(bounds, edge)
LOOP
IF maxbounds ~ g THEN
    RETURN NEXT g;
END IF;
END LOOP;
END;
$$
LANGUAGE 'plpgsql'
IMMUTABLE
STRICT
PARALLEL SAFE;


WITH c AS (select
2 as z,
0 as x,
1 as y),
d AS (
SELECT st_transform(g, 4326) geom FROM c, TileHexagons(c.z, c.x, c.y, 3) g
UNION
SELECT st_transform(ST_TileEnvelope(z, x, y), 4326) geom FROM c
)
SELECT ST_asgeojson(st_collect(geom)) FROM d;


select st_asgeojson(st_collect(st_transform(g, 4326))) from c, TileHexagons(3, 4, 4, 2) g;




```
CREATE FUNCTION zxy_houses(z integer, x integer, y integer, height float8)
RETURNS bytea
AS $$
DECLARE
rslt bytea;
BEGIN
  rslt := '123'::bytea;
  RETURN rslt;
END;
$$
LANGUAGE 'plpgsql'
;
```

```
CREATE OR REPLACE FUNCTION zxy_houses(z integer, x integer, y integer, OUT xy integer, OUT yz integer)
RETURNS SETOF record
AS $$
BEGIN
  FOR xy, yz IN SELECT a+x AS xy, a+y AS yz FROM generate_series(1,5) AS a
  LOOP
      RETURN NEXT;
  END LOOP;
END;
$$
LANGUAGE 'plpgsql';
```


CREATE FUNCTION zxy2_houses(z integer, x integer, y integer, height float8, OUT mvt bytea)
RETURNS bytea
AS $$
BEGIN
  mvt := '123'::bytea;
  RETURN;
END;
$$
LANGUAGE 'plpgsql'



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
