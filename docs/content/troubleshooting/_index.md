---
title: "Troubleshooting"
date:
draft: false
weight: 50
---

# Troubleshooting (optional)

## Tile Server

To get more information about what is going on behind the scenes, run with the `--debug` commandline parameter on, or turn on debugging in the configuration file:
```sh
./pg_tileserv --debug
```

## Web Layer

Hitting your service end points with a command-line utility like [curl](https://curl.haxx.se/) can also yield useful information:
```sh
curl -I http://localhost:7800/index.json
```

## Database Layer

The debug mode of the tile server returns the SQL that is being called on the database, but if you want to delve more deeply into all the SQL that is being run on the database, you can turn on [statement logging](https://www.postgresql.org/docs/current/runtime-config-logging.html#GUC-LOG-STATEMENT) in PostgreSQL by editing the `postgresql.conf` file for your database and restarting.

## Bug Reporting

If you find an issue with the tile server, bugs can be reported on GitHub at the issue tracker:

* https://github.com/crunchydata/pg_tileserv/issues
