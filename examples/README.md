# Web Map Tile Examples

* [Openlayers](./openlayers/openlayers-tiles.html)
* [Leaflet](./leaflet/leaflet-tiles.html)
* [Mapbox GL JS](./mapbox-gl-js/mapbox-gl-js-tiles.html)

## Data

Load the Natural Earth [Admin 0](https://www.naturalearthdata.com/downloads/50m-cultural-vectors/) boundaries into a table named `public.ne_50m_admin_0_countries` as the examples will attempt to load vectors from http://localhost:7800/public.ne_50m_admin_0_countries/{z}/{x}/{y}.pbf

```bash
shp2pgsql -D -s 4326 ne_50m_admin_0_countries.shp | psql -d naturalearth
```


