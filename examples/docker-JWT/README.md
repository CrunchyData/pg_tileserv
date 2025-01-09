# Docker Examples

This is a modified version of the [docker example](../docker), which tests JWT role switching.

It populates the database with a publicly-accessible table, and a private table that cna only be viewed if the appropriate `Authorization` header is provided.

This example uses Docker Compose with `pg_tileserv` and PostGIS (v3) Docker Images.

We use a [docker-compose.yml file](docker-compose.yml) with environment settings in
[pg_tileserv.env](pg_tileserv.env) and [pg.env](pg.env) for `pg_tileserv` and the PG database.


## The JWT setup

The test use a valid token:
 `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicm9sZSI6ImF1dGhvcml6ZWRfdXNlciIsImlhdCI6MTUxNjIzOTAyMn0.vU5_SXLaVOg5gjase2P2SJlsoo-oa7p_wWq0OQCD_9Q`

This is signed using the secret stored in the `TS_JWTSECRET` environment variable set in `pg_tileserv.env` (this is equivalent ot setting a value of `JwtSecret` in the config file), and has the payload

```json
{
  "sub": "1234567890",
  "role": "authorized_user",
  "iat": 1516239022
}
```

As a JWT Secret has been set, role-switching is enabled; and as the `JwtRoleClaimKey` has been left at tis default value of `role`, `pg_tileserv` will switch to the database role named `authorized_user` before executing queries in response to requests including this header.



## Build

* `docker-compose build`

This should build the latest [Alpine-based Docker Image](../../Dockerfile.alpine) for `pg_tileserv`.

## Run

* `docker-compose up`

N.B. on the first run the PostGIS Docker Image is downloaded and the DB initialized. `pg_tileserv` may not be able to connect.
In that case stop the Docker Compose process (ctrl-C) and run again. In another terminal window test
if `pg_tileserv` Container is at least running:

* `curl -v http://localhost:7800/jwt_test.locations_public/5/15/10.pbf`.

You will see an regular error message like *"Unable to get layer 'public.ne_50m_admin_0_countries"* as no data is yet in the database.


## Load Data

The file `test_setup.sql` populates the test database.
It is automatically run when the PostGIS docker container is run (when it is started subsequently it is ignored, unless the `docker-jwt_pg_tileserv_db` docker volume that persists the database data has been deleted).


## Run Webviewers

As the `pg_tileserv` container has a Docker port-mapping to localhost:7800, you can use the standard HTML examples locally in your browser.
The index pages at `http://localhost:7800/index.html` and  `http://localhost:7800/index.json` should list only the public layer, not the private layer. 

## Running tests

The file `test.hurl` contains a numebr of tests written in the format used by [Hurl](https://hurl.dev/).

If you have installed hurl, you can run these with `hurl -test test.hurl`.

You can also manually make equivalent requests using curl, with a command like:

  
    curl -v http://localhost:7800/jwt_test.locations_private/5/15/10.pbf -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicm9sZSI6ImF1dGhvcml6ZWRfdXNlciIsImlhdCI6MTUxNjIzOTAyMn0.vU5_SXLaVOg5gjase2P2SJlsoo-oa7p_wWq0OQCD_9Q"


## Clean/Restart

If something goes wrong along the way, or you want a clean restart, run this script:

* `./cleanup.sh`

This will delete dangling Docker Containers,  Images and the DB volume
