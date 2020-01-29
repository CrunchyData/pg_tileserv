---
title: "Introduction"
date:
draft: false
weight: 10
---

# pg_tileserv

`pg_tileserv` is a [PostGIS](https://postgis.net/)-only tile server in [Go](https://golang.org/).
 Strip away all the other requirements, it just has to take in HTTP tile request
s and form and execute SQL.  In a sincere act of flattery, the API and design look a lot like that of the [Martin](https://github.com/urbica/martin) tile server.

# Introduction

There are already lots of tile generators out there ([Tegola](https://tegola.io/), [Geoserver](https://geoserver.org), [Mapserver](https://mapserver.org)) that read from multiple data sources and generate vector tiles. In exchange for that flexibility of format, they provide less flexibility of usage. By restricting itself to only using PostGIS as a data source, `pg_tileserv` gains the following features:

* **Automatic configuration.** Just point the server at a PostgreSQL / PostGIS database, and the server can discover and automatically publish as tiles sources all tables it has read access to.
* **Full SQL flexibility.** Using [function layers]() the server can run any SQL at all to generate tile outputs. Any data processing or feature filtering or record aggregation you can express in SQL, you can expose as parameterized tile sources.
* **Database security model.** You can restrict access to tables and functions using standard database access control. This means you can also use advanced access control techniques, like row-level security to dynamically filter access based on the login role.

# Architecture

`pg_tileserv` is one component in "PostGIS for the Web" (aka "PostGIS FTW"), a growing family of Go spatial micro-services. Database-centric applications naturally have a central source of coordinating state, the database, which allows otherwise independent micro-services to coordinate and provide HTTP-level access to the database with relatively little middle-ware software complexity.

* [pg_tileserv](.) provides MVT tiles for interactive clients and smooth rendering
* [pg_featureserv]() provides GeoJSON feature services for reading and writing vector and attribute data from tables
* [pg_importserv]() (TBD) will provide an import API for ingesting arbitrary GIS data files

It should be possible to stand up a spatial services architecture of stateless microservices surrounding a PostgreSQL/PostGIS database cluster, in a standard container environment, on any cloud platform or internal datacenter: that's "PostGIS for the Web".

# Definitions

* **Map tiles** are a way of representing a multi-scale, [zoomable cartographic map](https://en.wikipedia.org/wiki/Tiled_web_map) by regularly subidividing the plane into independent tiles that can then rendered on a server and retrieved by a map client in parallel.
* **Vector tiles** are a [specific format of map tile](https://docs.mapbox.com/vector-tiles/specification/) that encode the features as vectors and delegate rendering of the features into cartography to the client web browser. Client side vector rendering uses less bandwidth, which is good for mobile clients, and allow more options for client side dynamic data visualizations.
* **Spatial database** is a database that includes a "geometry" column type. The PostGIS extension to PostgreSQL adds a geometry column type, and hundreds of functions to operate on that type, including the [ST_AsMVT()](https://postgis.net/docs/ST_AsMVT.html) function that `pg_tileserv` depends upon.

