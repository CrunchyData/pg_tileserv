# pg_tileserv

An experiment in a [PostGIS](https://postgis.net/)-only tile server in [Go](https://golang.org/). Strip away all the other requirements, it just has to take in HTTP tile requests and form and execute SQL.  In a sincere act of flattery, I have mostly copied the API of the [Martin](https://github.com/urbica/martin) tile server.

## Table Sources



## Function Sources

```
CREATE FUNCTION zxy_houses(z integer, x integer, y integer, height float8)
RETURNS bytea
AS $$
DECLARE
rslt bytea;
BEGIN
  rslt := '123'::bytea;
  RETURN rslt;
END;
$$
LANGUAGE 'plpgsql'
;
```

```
CREATE OR REPLACE FUNCTION zxy_houses(z integer, x integer, y integer, OUT xy integer, OUT yz integer)
RETURNS SETOF record
AS $$
BEGIN
  FOR xy, yz IN SELECT a+x AS xy, a+y AS yz FROM generate_series(1,5) AS a
  LOOP
      RETURN NEXT;
  END LOOP;
END;
$$
LANGUAGE 'plpgsql';
```


CREATE FUNCTION zxy2_houses(z integer, x integer, y integer, height float8, OUT mvt bytea)
RETURNS bytea
AS $$
BEGIN
  mvt := '123'::bytea;
  RETURN;
END;
$$
LANGUAGE 'plpgsql'



SELECT
Format('%s.%s', n.nspname, p.proname) AS id,
n.nspname,
p.proname,
d.description,
p.proargnames AS argnames,
string_to_array(oidvectortypes(p.proargtypes),', ') AS argtypes
FROM pg_proc p
JOIN pg_namespace n ON (p.pronamespace = n.oid)
LEFT JOIN pg_description d ON (p.oid = d.objoid)
WHERE p.proargtypes[0:2] = ARRAY[23::oid, 23::oid, 23::oid]
AND p.proargnames[1:3] = ARRAY['z'::text, 'x'::text, 'y'::text]
AND prorettype = 17
AND has_function_privilege(Format('%s.%s(%s)', n.nspname, p.proname, oidvectortypes(proargtypes)), 'execute') ;
