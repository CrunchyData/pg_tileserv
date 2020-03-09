CREATE OR REPLACE
-- Tile generating function takes in tile coordinates
-- and parameters and returns a tile in a bytea return.
-- We take in a click point and a count of how many
-- hydrants to build our voronoi with.
FUNCTION public.hydrants_voronoi(
            z integer, x integer, y integer,
            lon float8 default -123.129,
            lat float8 default 49.253,
            count bigint default 0)
RETURNS bytea
AS $$
  -- Find the N hydrants closest to our click point.
  -- The click point is geography coordinates (4326) and
  -- the hydrants are in UTM10 (26910) so we transform
  -- the click and run a nearest neighbor search.
  WITH hydrants_near AS (
    SELECT *
    FROM hydrants
    ORDER BY geom <-> ST_Transform(ST_SetSRID(ST_MakePoint(lon, lat),4326),26910)
    LIMIT count
  ),
  -- Convert the tile coordinates to an actual box
  -- web mercator.
  bounds AS (
    SELECT ST_TileEnvelope(z, x, y) AS geom
  ),
  -- Convert our hydrants into a collection (to build the vonoroi)
  -- and a hull (to clip the final output)
  hydrant_collection AS (
    SELECT
      ST_Collect(geom) AS collection,
      ST_Buffer(ST_ConvexHull(ST_Collect(geom)),50) AS shape
    FROM hydrants_near
  ),
  -- Build the voronoi, but only for tiles that intersect
  -- the clipping boundary of our result, this saves
  -- cycles as we zoom out.
  voronoi AS (
    SELECT (ST_Dump(ST_VoronoiPolygons(collection))).geom AS geom
    FROM hydrant_collection, bounds
    WHERE ST_Intersects(ST_Transform(bounds.geom, 26910), shape)
  ),
  -- Spatially join the polygons back to the hydrants,
  -- so we can add the hydrant attributes back onto the
  -- poygons. While we are here, also clip the voronoi polygons
  -- to our clipping shape.
  joined AS (
    SELECT ST_Intersection(v.geom, c.shape) AS vgeom, h.*
    FROM hydrants_near h
    JOIN voronoi v ON ST_Contains(v.geom, h.geom)
    CROSS JOIN hydrant_collection c
  ),
  -- Convert the clipped polygons and a selection of
  -- columns into the final MVT format for return.
  mvtgeom AS (
    SELECT ST_AsMVTGeom(ST_Transform(j.vgeom, 3857), bounds.geom) AS geom,
      j.status, j.color, j.code, j.subsystem, j.oocdate, j.oocnotes,
      j.street_numb, j.street
    FROM joined j, bounds
    WHERE ST_Intersects(j.vgeom, ST_Transform(bounds.geom, 26910))
  )
  SELECT ST_AsMVT(mvtgeom, 'public.hydrants_voronoi') FROM mvtgeom
$$
LANGUAGE 'sql'
STABLE
PARALLEL SAFE;

COMMENT ON FUNCTION public.hydrants_voronoi IS 'Given a tile address, click coordinates, and a feature count, generate a Voronoi diagram around the N nearest fire hydrants to the click point.';
