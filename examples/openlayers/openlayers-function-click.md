# Dynamic Voronoi Example

Start by downloading the "[fire hydrant data](https://opendata.vancouver.ca/explore/dataset/water-hydrants/download/?format=shp&timezone=America/Los_Angeles&lang=en&epsg=26910)" as a shape file from the City of Vancouver open data site.

Load the data into your PostgreSQL/PostGIS database using shp2pgsql:

```bash
shp2pgsql -s 26910 -D -I water-hydrants.shp hydrants | psql postgisftw
```

Create the `public.hydrants_delaunay()` function in your database by loading the [openlayer-function-click.sql](openlayer-function-click.sql) file.

Restart your `pg_tileserv` or at least reload the front page to force a refresh of the catalogue information.

Load the [openlayer-function-click.html](openlayer-function-click.html) file in your browser.
