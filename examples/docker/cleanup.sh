#!/bin/bash
# General cleanup, first stop

# first make sure the compose stack is stopped
docker-compose stop

#
# Remove all exited containers
for c in $(docker ps -a -f status=exited -q)
do
	docker rm ${c}
done

# And dangling images
for i in $(docker images -f dangling=true -q)
do
	docker rmi ${i}
done

# Remove the DB Volume, removes database
docker volume rm docker_pg_tileserv_db
