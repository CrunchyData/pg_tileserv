#!/bin/bash

DOCKER_REPO=pramsey/pg_tileserv


if [ "$TARGET" = "windows" ]; then
    BINARY=pg_tileserv.exe
else
    BINARY=pg_tileserv
fi

if [ "$TRAVIS_TAG" = "" ]; then
    TAG=latest
else
    TAG=$TRAVIS_TAG
fi

# docker deploy
if [ "$TARGET" = "docker" ]; then
    DATE=`date +%Y%m%d`
    make build-docker
    docker tag $DOCKER_REPO:$TAG $DOCKER_REPO:$DATE
    if [ "$TRAVIS_BRANCH" = "master" ] && [ "$TRAVIS_PULL_REQUEST" = "false" ]; then
        docker login -u "$DOCKER_USER" -p "$DOCKER_PASS"
        docker push $DOCKER_REPO
    fi
# windows, linux, osx pre-deploy
elif [ "$TARGET" != "docs" ]; then
    mkdir upload
    zip -r upload/pg_tileserv_${TAG}_${TARGET}.zip ${BINARY} README.md LICENSE.md assets/ config/
fi
