---
title: "Examples"
date:
draft: false
weight: 30
---

The web map examples in this section are set up to render a basemap layer from Wikimedia and vector tiles from `pg_tileserver` 
running on a local machine, using popular open source JavaScript web map components.

## Load Natural Earth Data

### Database preparation

The following terminal commands will create a database named `naturalearth`, assuming that your user account has create database privilege:

```
createdb naturalearth
```

Load the PostGIS extension as superuser (`postgres`):

```
psql -U postgres -d naturalearth -c 'CREATE EXTENSION postgis'
```

### Import shapefile

The data used in the examples are loaded from [Natural Earth](https://www.naturalearthdata.com/downloads/50m-cultural-vectors/). Download the *Admin 0 - Countries* ZIP and extract to a location on your 
machine. 

In that directory, run the following command in the terminal to load the 
shapefile data into the `naturalearth` database. This creates a new table `ne_50m_admin_0_countries`, with the application user as the owner -- refer to [Table Layers](../usage/table-layers/) and [Security](../usage/security/) for more information on access to spatial tables on `pg_tileserv`.

```
shp2pgsql -D -s 4326 ne_50m_admin_0_countries.shp | psql -U username -d naturalearth
```

You should see the `ne_50m_admin_0_countries` table with the `\dt` SQL shell command.

Make sure that `pg_tileserv` connection specifies `naturalearth`, i.e.: `DATABASE_URL=postgres://username:password@host/naturalearth`. With the service running, you should also see the layer on the web preview, i.e.: http://localhost:7800/public.ne_50m_admin_0_countries.html 

![pg_tileserv web interface preview](/example-web-preview.png)
