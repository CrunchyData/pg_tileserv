CREATE SCHEMA jwt_test;

CREATE TABLE jwt_test.locations_private (
	name varchar(254) NULL,
	geom public.geometry(geometry, 27700) NULL
);
INSERT INTO jwt_test.locations_private (name, geom) VALUES ('1', ST_GeomFromText('POINT(531874 190299)', 27700));
INSERT INTO jwt_test.locations_private (name, geom) VALUES ('2', ST_GeomFromText('POINT(532058 167349)', 27700));
INSERT INTO jwt_test.locations_private (name, geom) VALUES ('3', ST_GeomFromText('POINT(541745 185015)', 27700));
INSERT INTO jwt_test.locations_private (name, geom) VALUES ('4', ST_GeomFromText('POINT(531272.97 193512.79)', 27700));
INSERT INTO jwt_test.locations_private (name, geom) VALUES ('5', ST_GeomFromText('POINT(534947 185980)', 27700));


CREATE TABLE jwt_test.locations_public (
	name varchar(254) NULL,
	geom public.geometry(geometry, 27700) NULL
);
INSERT INTO jwt_test.locations_public (name, geom) VALUES ('1', ST_GeomFromText('POINT(531874 190299)', 27700));
INSERT INTO jwt_test.locations_public (name, geom) VALUES ('2', ST_GeomFromText('POINT(532058 167349)', 27700));
INSERT INTO jwt_test.locations_public (name, geom) VALUES ('3', ST_GeomFromText('POINT(541745 185015)', 27700));
INSERT INTO jwt_test.locations_public (name, geom) VALUES ('4', ST_GeomFromText('POINT(531272.97 193512.79)', 27700));
INSERT INTO jwt_test.locations_public (name, geom) VALUES ('5', ST_GeomFromText('POINT(534947 185980)', 27700));


-- a user that exists but can't be switched to
create user cant_switch_user;

-- a user with no access
create user no_access_user;

-- an anonymous user that has access to locations_public but not locations_private
create user anonymous_user;
grant usage on schema jwt_test to anonymous_user ;
grant select on jwt_test.locations_public to anonymous_user;


-- an authorized user that has access to both locations_public and locations_private
create user authorized_user;
grant usage on schema jwt_test to  authorized_user;
grant select on jwt_test.locations_public to authorized_user;
grant select on jwt_test.locations_private to authorized_user;


-- the authenticator user, that has ability to switch role to any of anonymous_user/authorized_user/no_access_user
create user authenticator_user noinherit;
ALTER USER authenticator_user WITH PASSWORD 'the_authenticator_password';
grant anonymous_user to authenticator_user;
grant no_access_user to authenticator_user;
grant authorized_user to authenticator_user;
