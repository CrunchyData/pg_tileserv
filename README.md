# pg_tileserv

An experiment in a [PostGIS](https://postgis.net/)-only tile server in [Go](https://golang.org/). Strip away all the other requirements, it just has to take in HTTP tile requests and form and execute SQL.  In a sincere act of flattery, I have mostly copied the API of the [Martin](https://github.com/urbica/martin) tile server.

## Table Sources



## Function Sources

CREATE FUNCTION xyz_houses(x integer, y integer, z integer, height float8)
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


SELECT
  p.oid, proname,
  string_to_array(oidvectortypes(proargtypes),', ') AS argtypes,
  proargnames AS argnames,
  d.description,
  n.nspname
FROM pg_proc p
JOIN pg_namespace n ON (p.pronamespace = n.oid)
LEFT JOIN pg_description d ON (p.oid = d.objoid)
WHERE proname LIKE 'xyz_%'
AND proargtypes[0:2] = ARRAY[23::oid, 23::oid, 23::oid]
AND prorettype = 17;
