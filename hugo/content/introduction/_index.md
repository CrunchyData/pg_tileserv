---
title: "About pg_tileserv"
date:
draft: false
weight: 10
---

## Motivation

There are numerous tile generators available (such as [Tegola](https://tegola.io/), [Geoserver](https://geoserver.org), [Mapserver](https://mapserver.org)) that read from multiple data sources and generate vector tiles. `pg_tileserv` works exclusively with PostGIS data, but this also allows more flexibility of usage.

By restricting itself to only using PostGIS as a data source, `pg_tileserv` gains the following features:

* **Automatic configuration.** The server can discover and automatically publish as tiles sources all tables it has read access to: just point it at a PostgreSQL/PostGIS database.
* **Full SQL flexibility.** Using [function layers](/usage/function-layers/), the server can run any SQL to generate tile outputs. Any data processing, feature filtering, or record aggregation that can be expressed in SQL, can be exposed as parameterized tile sources.
* **Database security model.** You can restrict access to tables and functions using standard database access control. This means you can also use advanced access control techniques, like row-level security to dynamically filter access based on the login role.

## Architecture

`pg_tileserv` is one component in "PostGIS for the Web" (aka "PostGIS FTW"), a growing family of Go spatial microservices. Database-centric applications naturally have a central source of coordinating state, the database, which allows otherwise independent microservices to coordinate and provide HTTP-level access to the database with less middleware software complexity.

* [pg_tileserv](/) provides MVT tiles for interactive clients and smooth rendering
* [pg_featureserv](https://access.crunchydata.com/documentation/pg_featureserv/latest/) provides GeoJSON feature services for reading and writing vector and attribute data from tables

PostGIS for the Web makes it possible to stand up a spatial services architecture of stateless microservices surrounding a PostgreSQL/PostGIS database cluster, in a standard container environment, on any cloud platform or internal datacenter.

## Definitions

* **Map tiles** are a way of representing a multi-scale, [zoomable cartographic map](https://en.wikipedia.org/wiki/Tiled_web_map) by regularly subidividing the plane into independent tiles that can then be rendered on a server and retrieved by a map client in parallel.
* **Vector tiles** are a [specific format of map tile](https://docs.mapbox.com/vector-tiles/specification/) that encode the features as vectors and delegate to the client web browser the rendering of the features into cartography. Client side vector rendering uses less bandwidth, which is good for mobile clients, and allow more options for client side dynamic data visualizations.
* A **spatial database** is a database that includes a "geometry" column type. The PostGIS extension to PostgreSQL adds a geometry column type, as well as hundreds of functions to operate on that type, including the [ST_AsMVT()](https://postgis.net/docs/ST_AsMVT.html) function that `pg_tileserv` depends upon.
