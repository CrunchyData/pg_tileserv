---
title: "Installation"
date:
draft: false
weight: 20
---


## Requirements

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

## Installation

To install `pg_tileserv`, download the binary file. Alternatively, you may run a container, or build the executable from source.

### A. Download Binaries

Builds of the latest code:

* [Linux](https://postgisftw.s3.amazonaws.com/pg_tileserv_latest_linux.zip)
* [Windows](https://postgisftw.s3.amazonaws.com/pg_tileserv_latest_windows.zip)
* [OSX](https://postgisftw.s3.amazonaws.com/pg_tileserv_latest_osx.zip)

Unzip the file, copy the `pg_tileserv` binary wherever you wish, or use it in place. If you move the binary, remember to move the `assets/` directory to the same location, or start the server using the `AssetsDir` configuration option.

### B. Run Container

There is a docker image available on DockerHub.

* [Docker](https://hub.docker.com/repository/docker/pramsey/pg_tileserv)

Run the container, provide database connection information in the `DATABASE_URL` environment variable and map the default service port (7800).

```sh
docker run -e DATABASE_URL=postgres://user:pass@host/dbname -p 7800:7800 pramsey/pg_tileserv
```

### C. Build From Source

Install the [Go software development environment](https://golang.org/doc/install). Make sure that the [`GOPATH` environment variable](https://github.com/golang/go/wiki/SettingGOPATH) is also set.

```sh
SRC=$GOPATH/src/github.com/CrunchyData
mkdir -p $SRC
cd $SRC
git clone git@github.com:CrunchyData/pg_tileserv.git
cd pg_tileserv
go build
go install
```

To run the build, set the `DATABASE_URL` environment variable to the database you want to connect to, and run the binary.

```sh
export DATABASE_URL=postgres://user:pass@host/dbname
$GOPATH/bin/pg_tileserv
```

## Deployment

### Basic Operation

#### Linux/OSX

```sh
export DATABASE_URL=postgresql://username:password@host/dbname
./pg_tileserv
```

#### Windows

```
SET DATABASE_URL=postgresql://username:password@host/dbname
pg_tileserv.exe
```

### Configuration File

The configuration file will be automatically read from the following locations, if it exists:

* In the system configuration directory, at `/etc/pg_tileserv.toml`
* Relative to the directory from which the program is run, `./pg_tileserv.toml`

If you want to pass a path directly to the configuration file, use the `--config` commandline parameter to pass in a pull path to configuration file. When using the `--config` option, configuration files in other locations will be ignored.

```sh
./pg_tileserv --config /opt/pg_tileserv/pg_tileserv.toml
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
