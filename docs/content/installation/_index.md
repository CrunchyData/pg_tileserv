---
title: "Installation"
date:
draft: false
weight: 20
---

This is just a short little page saying - hey we are about to cover installation

# Requirements

* **PostgreSQL 9.5** or later
* **PostGIS 2.4** or later

The tile server depends on the [ST_AsMVT()](https://postgis.net/docs/ST_AsMVT.html) function, which is only available if PostGIS has been compiled with support for the **libprotobuf** library. See the output from [PostGIS_Full_Version](https://postgis.net/docs/PostGIS_Full_Version.html), for example:

```sql
SELECT postgis_full_version()
```
```
POSTGIS="3.0.1" [EXTENSION] PGSQL="121" GEOS="3.8.0-CAPI-1.13.1 "
PROJ="6.1.0" LIBXML="2.9.4" LIBJSON="0.13"
LIBPROTOBUF="1.3.2" WAGYU="0.4.3 (Internal)"
```

# Installation

## Download Binaries

Builds of the latest code:

* [Linux](https://postgisftw.s3.amazonaws.com/pg_tileserv_latest_linux.zip)
* [Windows](https://postgisftw.s3.amazonaws.com/pg_tileserv_latest_windows.zip)
* [OSX](https://postgisftw.s3.amazonaws.com/pg_tileserv_latest_osx.zip)

Unzip the file, copy the `pg_tileserv` binary wherever you wish, or use it in place. If you move the binary, remember to move the `assets/` directory to the same location, or start the server using the `AssetsDir` configuration option.

## Container

There is a docker image available on DockerHub.

* [Docker](https://hub.docker.com/repository/docker/pramsey/pg_tileserv)

Run the container, provide database connection information in the `DATABASE_URL` and map the default service port (7800).

```sh
docker run -e DATABASE_URL=postgres://user:pass@host/dbname -p 7800:7800 pramsey/pg_tileserv
```

## Build From Source

Install the [Go software development environment](https://golang.org/doc/install).

```sh
SRC=$GOPATH/src/github.com/CrunchyData
mkdir -p $SRC
cd $SRC
git clone git@github.com:CrunchyData/pg_tileserv.git
cd pg_tileserv
go build
go install
```

To run the build, set the `DATABASE_URL` to the database you want to connect to, and run the binary.

```sh
export DATABASE_URL=postgres://user:pass@host/dbname
$GOPATH/bin/pg_tileserv
```

# Deployment

## Basic Operation

### Linux/OSX

```sh
export DATABASE_URL=postgresql://username:password@host/dbname
./pg_tileserv
```

### Windows

```
SET DATABASE_URL=postgresql://username:password@host/dbname
pg_tileserv.exe
```

## Trouble-shooting

To get more information about what is going on behind the scenes, run with the `--debug` commandline parameter on, or turn on debugging in the configuration file:
```sh
./pg_tileserv --debug
```

## Configuration File

If you want to alter default values other than the database connection, use the `--config` commandline parameter to pass in a configuration file.

```sh
./pg_tileserv --config /etc/pg_tileserv.toml
```

In general the defaults are fine, and the program autodetects things like the server name.

```toml
# Database connection
DbConnection = "user=you host=localhost dbname=yourdb"
# Close pooled connections after this interval
DbPoolMaxConnLifeTime = "1h"
# Hold no more than this number of connections in the database pool
DbPoolMaxConns = 4
# Look to read html templates from this directory
AssetsPath = "./assets"
# Accept connections on this subnet (default accepts on all subnets)
HttpHost = "0.0.0.0"
# Accept connections on this port
HttpPort = 7800
# Advertise URLs relative to this server name
# default is to look this up from incoming request headers
# UrlBase = "http://yourserver.com/"
# Resolution to quantize vector tiles to
DefaultResolution = 4096
# Rendering buffer to add to vector tiles
DefaultBuffer = 256
# Limit number of features requested (-1 = no limit)
MaxFeaturesPerTile = 10000
# Advertise this minimum zoom level
DefaultMinZoom = 0
# Advertise this maximum zoom level
DefaultMaxZoom = 22
# Output extra logging information?
Debug = false
```


