---
title: "Security"
date:
draft: false
weight: 500
---

The basic principle of security is to connect your tile server to the database with a user that has just the access you want it to have, and no more.

Start with a new, blank user. A blank user has no select privileges on tables it does not own.
It does have execute privileges on functions.
However, the user has no select privileges on tables accessed by functions, so effectively the user will still have no access to data.

```sql
CREATE USER tileserver;
```
To support different access patterns, create different users with access to different tables/functions, and run multiple services, connecting with those different users.

## Tables and Views

If your tables and views are in a schema other than `public`, you will have to also grant "usage" on that schema to your user.
```sql
GRANT USAGE ON SCHEMA myschema TO tileserver;
```
You can then grant access to the user one table at a time.
```sql
GRANT SELECT ON TABLE myschema.mytable TO tileserver;
```
Alternatively, you can grant access to all the tables at once.
```sql
GRANT SELECT ON ALL TABLES IN SCHEMA myschema TO tileserver;
```

## Functions

As noted above, functions that access table data are effectively restricted by the access levels the user has to the tables the function reads. If you want to completely restrict access to the function, including visibility in the user interface, you can strip execution privileges from the function.
```sql
-- All functions grant execute to 'public' and all roles are
-- part of the 'public' group, so public has to be removed
-- from the executors of the function
REVOKE EXECUTE ON FUNCTION myschema.myfunction FROM public;
-- Just to be sure, also revoke execute from the user
REVOKE EXECUTE ON FUNCTION myschema.myfunction FROM tileserver;
```
