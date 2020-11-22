# Docker Examples

by: Just van den Broecke - justb4 (@) gmail.com

This example uses Docker Compose with `pg_tileserv` and PostGIS (v3) Docker Images.
Run with these steps: Build, Run, Load Vector Data, and run the
standard web-viewer examples like [Leaflet](../leaflet/leaflet-tiles.html).

We use a [docker-compose.yml file](docker-compose.yml) with environment settings in
[pg_tileserv.env](pg_tileserv.env) and [pg.env](pg.env) for `pg_tileserv` and the PG database.

## Build

* `docker-compose build`

This should build the latest [Alpine-based Docker Image](../../Dockerfile.alpine) for `pg_tileserv`.

## Run

* `docker-compose up`

NB on the first run the PostGIS Docker Image is downloaded and the DB initialized. `pg_tileserv` may not be able to connect.
In that case stop the Docker Compose process (ctrl-C) and run again. In another terminal window test
if `pg_tileserv` Container is at least running:

* `curl -v http://localhost:7800/public.ne_50m_admin_0_countries/2/2/3.pbf`.

You will see an regular error message like *"Unable to get layer 'public.ne_50m_admin_0_countries"* as no data is yet in the database.

## Load Data

Load Natural Earth [Admin 0 Countries](https://www.naturalearthdata.com/downloads/50m-cultural-vectors/) boundaries into a table
named `public.ne_50m_admin_0_countries`. Download and unzip `ne_50m_admin_0_countries.zip`.

Our Postgres/PostGIS container is running on external port 5433 (which is mapped to internal Container standard PG port 5432)
hence we pipe `shp2pgsql` to `psql -U tileserv -p 5433 -h 0.0.0.0 -d tileserv`

* `shp2pgsql -D -s 4326 ne_50m_admin_0_countries.shp | psql -U tileserv -p 5433 -h 0.0.0.0 -d tileserv`

Provide the password (see [pg.env](pg.env) and restart Docker compose.

## Run Webviewers

As the `pg_tileserv` container has a Docker port-mapping to localhost:7800, you can use the standard webviewer examples.
like [Leaflet](../leaflet/leaflet-tiles.html), [MapBox](../mapbox-gl-js/mapbox-gl-js-tiles.html) and [OpenLayers](../openlayers/openlayers-tiles.html).

## Next
Run also the [OpenLayers Voronoi example](../openlayers/openlayers-function-click.md) but with Docker. This example demonstrates
the powerful "Function" capability of `pg_tileserv`.

Download "[fire hydrant data](https://opendata.vancouver.ca/explore/dataset/water-hydrants/download/?format=shp&timezone=America/Los_Angeles&lang=en&epsg=26910)"
as a shape file from the City of Vancouver open data site.

Unzip and load the data into the database like above for Countries:

* `shp2pgsql -s 26910 -D -I water-hydrants.shp hydrants | psql -U tileserv -p 5433 -h 0.0.0.0 -d tileserv`

Create the `public.hydrants_delaunay()` function in your database by loading the [openlayers-function-click.sql](../openlayers/openlayers-function-click.sql) file:

In this Docker example directory do:

* `cat ../openlayers/openlayers-function-click.sql | psql -U tileserv -p 5433 -h 0.0.0.0 -d tileserv`

Restart the Docker Compose process and load the [openlayers-function-click.html](../openlayers/openlayers-function-click.html) file in your browser.
