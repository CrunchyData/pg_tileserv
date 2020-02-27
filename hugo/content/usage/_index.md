---
title: "Usage"
date:
draft: false
weight: 25
---

The purpose of `pg_tileserv` is to turn a set of spatial records into tiles, on the fly. The tile server reads two different layers of data:

* **Table layers** are what they sound like: tables and views in the database that have a spatial column with a spatial reference system defined on it.
* **Function layers** hide the source of data from the server, and allow the HTTP client to send in optional parameters to allow more complex SQL functionality. Any function of the form `function(z integer, x integer, y integer, ...)` that returns an MVT `bytea` result can serve as a function layer.
