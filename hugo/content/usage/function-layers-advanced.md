---
title: "Advanced Function Layers"
date:
draft: false
weight: 400
---

## Dynamic Geometry Example

So far, all our examples have used simple SQL functions, but using the procedural [PL/pgSQL language](https://www.postgresql.org/docs/current/plpgsql.html) we can create much more interactive examples.

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
## Dynamic Hexagons with Spatial join Example

Hexagonal tilings are popular with data visualization experts because they can be used to summarize point data without adding a visual bias to the output via different summary area sizes. They also have a nice "non-pointy" shape, while still providing a complete tiling of the plane.

When you want to provide a hexagonal summary of a data set at multiple scales, it presents an implementation problem: do you need to create a pile of hexagon tables, solely for the purpose of summary visualization?

The answer is no: you can generate your hexagons dynamically based on the scale of the requested map tiles.

### Generate hexagons

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

### Find hexagon coordinates within the map tile

Now we need a function that, given a square input (a map tile), can figure out all the hexagon coordinates that fall within the tile. Again, the edge size of the hexagon tiling determines the overall geometry of the hex tiling. More than one hexagon will be required most times, so this is a set-returning function.
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

### Generate hexagons that cover the map tile

Next, we need a function that puts the two parts together: with tile coordinates and edge size as input, generate the set of all the hexagons that cover the tile. The output here is basically a spatial table: a set of rows, each row containing a geometry (hexagon) and some properties (hexagon coordinates). This is the input we need for a spatial join.
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
The function that the tile server actually calls looks like all other tile server functions: tile coordinates and optional parameter as input, `bytea` MVT as output.
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
SELECT ST_AsMVT(mvt, 'public.hexpopulationsummary') FROM mvt
$$
LANGUAGE 'sql'
STABLE
STRICT
PARALLEL SAFE;

COMMENT ON FUNCTION public.hexpopulationsummary IS 'Hex summary of the ne_50m_populated_places table. Step parameter determines how approximately many hexes (2^step) to generate per tile.';
```
