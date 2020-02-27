---
title: "About Tiles"
date:
draft: false
weight: 600
---

The tiles produced by PostGIS and published via `pg_tileserv` are "[Mapbox vector tiles](https://github.com/mapbox/vector-tile-spec)", a widely used de facto standard encoding of vector tiles.

The purpose of vector tiles is to efficiently transfer map features over the network, so they optimize for size, using a variety of techniques while retaining enough context to be useful to the client mapping environment.

## Resolution

Coordinates in tiles are quantized to integer values, and the default resolution of vector tiles is 4096 by 4096. The default resolution can be altered using the `DefaultResolution` [configuration parameter](/installation#configuration-file/).

## Tile Buffer

Tiles are rendered independently. For features with wide styles near borders, a copy of the feature needs to appear in both neighboring tiles, or a rendering failure will occur.

![Tile rendering failure](/tile-render-failure.png)

The default tile buffer is 256 pixels, which is enough for most rendering cases. You can make your tiles smaller if you have narrow rendering styles, by reducing the `DefaultBuffer` configuration parameter.

## Unique Identifier

The [vector tile specification](https://github.com/mapbox/vector-tile-spec) includes an optional "[id](https://github.com/mapbox/vector-tile-spec/blob/master/1.0.1/vector_tile.proto#L30)" element that provides a unique feature identifier.

A single feature can end up in multiple tiles, and the unique identifier allows the client side renderer to do things like roll-overs and highlights on features that span tile boundaries.

The tile server can automatically populate the "id" element, but only in cases where:

* PostGIS version is >= 3.0, as the [ST_AsMVT()](https://postgis.net/docs/ST_AsMVT.html) function did not support feature id until then.
* The table being published has a integer primary key defined. This key will be used as the "id" automatically.

For function layers, the "id" can be populated, but that task is left to the function author, who will be calling the `ST_AsMVT()` function in their code, and must remember to populate the feature id name field. The column that is chosen to populate the "id" element **must** be unique per feature.
