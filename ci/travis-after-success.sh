#!/bin/bash

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

if [ "$TARGET" = "docker" ]; then
    DOCKER_REPO=pramsey/pg_tileserv
    VERSION=`./pg_tileserv --version | cut -f2 -d' '`
    DATE=`date +%Y%m%d`
    docker login -u "$DOCKER_USER" -p "$DOCKER_PASS"
    docker build -f Dockerfile --build-arg VERSION=$VERSION -t $DOCKER_REPO:$TAG .
    #docker tag $DOCKER_REPO:$TAG $DOCKER_REPO:$TRAVIS_COMMIT
    #docker tag $DOCKER_REPO:$TAG $DOCKER_REPO:travis-$TRAVIS_BUILD_NUMBER
    docker tag $DOCKER_REPO:$TAG $DOCKER_REPO:$DATE
    docker push $DOCKER_REPO
else
    mkdir upload
    zip -r upload/pg_tileserv_${TAG}_${TARGET}.zip ${BINARY} README.md LICENSE.md assets/
fi
