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

We load all sample data using `shp2pgsql` and `psql` within the PostGIS Docker Container, so we don't need to install any Postgres/PostGIS tools locally.
The `./data` dir is mapped into the Docker Container at `/work`.

First Download these files into the `./data` subdir:

* Natural Earth [Admin 0 Countries](https://www.naturalearthdata.com/downloads/50m-cultural-vectors/).
* [fire hydrant data](https://opendata.vancouver.ca/explore/dataset/water-hydrants/download/?format=shp&timezone=America/Los_Angeles&lang=en&epsg=26910)"

Unzip these two zip-files within the `./data` subdir.

To run also the [OpenLayers Voronoi example](../openlayers/openlayers-function-click.md) using Docker, we apply
the [OpenLayers Function-click SQL](../openlayers/openlayers-function-click.sql). This example demonstrates the powerful "Function" capability of `pg_tileserv`,
creating the `public.hydrants_delaunay()` function in your database.

To load the two datasets and Function SQL, use the [load-data.sh helper script](load-data.sh)

* `./load-data.sh`
* restart the docker-compose stack

The above data-loading script `exec`s the running PostGIS Docker Container `pg_tileserv_db` as for example:

* `docker-compose exec pg_tileserv_db sh -c "shp2pgsql -d -D -s 4326 /work/ne_50m_admin_0_countries.shp | psql -U tileserv -d tileserv"`

## Run Webviewers

As the `pg_tileserv` container has a Docker port-mapping to localhost:7800, you can use the standard HTML examples locally in your browser.
In a real-world application you would run these in a web-server container like `nginx` or `Apache httpd`.

See [Leaflet](../leaflet/leaflet-tiles.html), [MapBox](../mapbox-gl-js/mapbox-gl-js-tiles.html) and [OpenLayers](../openlayers/openlayers-tiles.html).
And the  [openlayers-function-click.html](../openlayers/openlayers-function-click.html) for the Voronoi Function example.

## Clean/Restart

If something goes wrong along the way, or you want a clean restart, run this script:

* `./cleanup.sh`

This will delete dangling Docker Containers and and Images and the DB volume
