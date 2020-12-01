# Shortcuts to load all data in PostGIS DB in Docker Container
# NB the docker-compose stack must be running !
#

# Load Admin 0 countries
docker-compose exec pg_tileserv_db sh -c "shp2pgsql -D -s 4326 /work/ne_50m_admin_0_countries.shp | psql -U tileserv -d tileserv"

# Load Vancouver Water Hydrants
docker-compose exec pg_tileserv_db sh -c "shp2pgsql -D -s 26910 -I /work/water-hydrants.shp hydrants | psql -U tileserv -d tileserv"

# Load SQL Functions for OpenLayers example
cp ../openlayers/openlayers-function-click.sql ./data/
docker-compose exec pg_tileserv_db sh -c "cat /work/openlayers-function-click.sql | psql -U tileserv -d tileserv"
rm ./data/openlayers-function-click.sql
