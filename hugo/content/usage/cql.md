---
title: "CQL Filters"
date:
draft: false
weight: 250
---

The features returned by `items` queries can be filtered using
the `filter` query parameter with an expression written using
the [Common Query Language](https://portal.ogc.org/files/96288) (CQL).
CQL expressions return a value of `true` or `false`.
Only features which evaluate to `true` are returned.

The filter expression is evaluated by the database,
which will take advantage of indexes (attribute and spatial)
to make filter evaluation very efficient.

This section describes the CQL query language subset supported.
Note that some special characters may need to be [URL-encoded](https://en.wikipedia.org/wiki/Percent-encoding) in some clients.

## Property and Literal Values

The basic elements of filter expressions are values obtained
from feature collection properties, and literals.

Properties are referred to by name.
Property names can be quoted, to support including special characters.

```
propname
"quoted_name$"
```

Literals can be numbers, boolean or text values.
To include single quotes in text values use repeated single quotes.

```
1.234
true
'a text value'
'it''s easy'
```

## Arithmetic expressions

Numeric values can be combined using the
arithmetic operators `+`,`-`,`*`,`/` and `%` (modulo), with parentheses if needed
for operator precedence.

**NOTE:** `+` needs to be URL-encoded as `%2B`.
```
x + 3
2 * (y - 3)
p % 10
```

## Comparisons

Values can be compared using conditional operators:
```
a = b   a <> b   a > b   a >= b   a < b   a <= b
```

#### Examples
```
pop_est >= 1000000
name = 'Finland'
```

## BETWEEN predicate

The `BETWEEN` predicate tests if a value lies in a range defined by a start and end value (inclusive):
```
property [NOT] BETWEEN a AND b
```

#### Examples
```
pop_est BETWEEN 100000 AND 1000000
name NOT BETWEEN 'Chile' AND 'Denmark'
```

## IN predicate
The `IN` predicate tests if a value lies in a list of constant values.
```
property [NOT] IN ( val1, val2, ... )
```

#### Examples
```
id IN (1,2,3)
name IN ('Chile', 'Kenya', 'Denmark')
```

## LIKE predicate
The `LIKE` predicate tests if a text value matches a pattern.
The character `%` is a wildcard.
(Note that this may need to be URL-encoded as `%25`.)
`ILIKE` can be used for case-independent matching.

```
property [NOT] LIKE | ILIKE pattern
```

#### Examples
```
name LIKE 'Ch%'
continent ILIKE '%america'
```

## IS NULL predicate
The `IS NULL` predicate tests if a property value is (or is not) null.
```
property IS [NOT] NULL
```

#### Example
```
name IS NULL
```

## Boolean combinations
Comparisons and predicates can be combined with the
boolean operators `AND`, `OR` and `NOT`.
Operators are evaluated in the order NOT, AND, OR.
Evaluation order can be controlled by enclosing
subexpressions in parentheses.

#### Example
```
(continent = 'Europe' OR continent = 'Africa') AND pop_est < 1000000
```

# Spatial filters

CQL supports spatial filtering by providing **geometry literals**
and **spatial predicates**.

## Geometry Literals

Geometry literals use [Well-Known Text](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry)
(WKT) to describe
values for points, lines, polygons (with holes), and collections.

#### Examples
```
POINT (1 2)
LINESTRING (0 0, 1 1)
POLYGON ((0 0, 0 9, 9 0, 0 0))
POLYGON ((0 0, 0 9, 9 0, 0 0),(1 1, 1 8, 8 1, 1 1))
MULTIPOINT ((0 0), (0 9))
MULTILINESTRING ((0 0, 1 1),(1 1, 2 2))
MULTIPOLYGON (((1 4, 4 1, 1 1, 1 4)), ((1 9, 4 9, 1 6, 1 9)))
GEOMETRYCOLLECTION(POLYGON ((1 4, 4 1, 1 1, 1 4)), LINESTRING (3 3, 5 5), POINT (1 5))
```

CQL also provides a syntax for concisely representing a rectangular polygon
by the X and Y ordinates at the lower-left and upper-right corners:
```
ENVELOPE (1, 2, 3, 4)
```

By default the coordinate system of geometry literal values is assumed to be geodetic (SRID = 4326).
The `filter-crs=SRID` query parameter can be used to specify that the geometry literals in a filter expression are in a different coordinate system.

## Spatial predicates

Spatial predicates allow filtering features via spatial conditions
on the feature geometry.
Spatial predicates are defined in the form of spatial functions.
Predicates for spatial relationships include:

* `INTERSECTS` - tests whether two geometries intersect
* `DISJOINT` - tests whether two geometries have no points in common
* `CONTAINS` - tests whether a geometry contains another
* `WITHIN` - tests whether a geometry is within another
* `EQUALS` - tests whether two geometries are topologically equal
* `CROSSES` - tests whether the geometries cross
* `OVERLAPS` - tests whether the geometries overlap
* `TOUCHES` - tests whether the geometries touch

For detailed definitions of the spatial predicates see the
[CQL standard](https://portal.ogc.org/files/96288#enhanced-spatial-operators)
and the [PostGIS function reference](https://postgis.net/docs/reference.html#Spatial_Relationships).

Typically a spatial predicate is used to test the relationship between the spatial column of the queried collection
and a geometry literal value.

#### Examples
```
INTERSECTS(geom, ENVELOPE(-100, 49, -90, 50) )

CONTAINS(geom, POINT(-100 49) )
```

The `DWITHIN` predicate allows testing whether a geometry lies within a given distance of another.  The distance is in the units of the dataset's coordinate system
(degrees in the case of data stored in SRID=4326, or a length unit such as meters for non-geodetic data).

#### Example
```
DWITHIN(geom, POINT(-100 49), 0.1)
```

## Temporal filters

Temporal filtering in CQL supports date-time literals and the ability to use
them in conditions against temporal-valued properties
(PostgreSQL [dates or timestamps](https://www.postgresql.org/docs/current/datatype-datetime.html)).

Date-time literals specifiy a date, or a timestamp including a date and time (with optional seconds value):
```
2001-01-01
2001-01-01T10:23
2001-01-01T10:23:45
```

Temporal values can be compared using the conditional operators `<`,`<=`,`>`,`>=`,`=`,`<>`:

```
t > 2001-01-01T00:00 AND t <= 2002-12-31T11:59:59
```

They can also be used in the `BETWEEN` predicate:
```
t BETWEEN 2001-01-01 AND 2001-12-31
```
