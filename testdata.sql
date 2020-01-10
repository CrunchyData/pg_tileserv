CREATE EXTENSION postgis;

CREATE TABLE public.no_geometry (id integer PRIMARY KEY, name text);

CREATE TABLE public.geometry_no_srid (id integer PRIMARY KEY, geom Geometry(Point));

CREATE TABLE public.geometry_only (geom Geometry(Point, 4326));

CREATE TABLE public.geometry_no_type (id integer PRIMARY KEY, geom Geometry(Geometry, 4326));

CREATE TABLE public.geometry_no_data (id integer PRIMARY KEY, geom Geometry(Geometry, 4326));
