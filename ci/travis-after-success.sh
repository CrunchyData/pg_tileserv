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

# docker deploy
if [ "$TARGET" = "docker" ]; then
    DOCKER_REPO=pramsey/pg_tileserv
    VERSION=`./pg_tileserv --version | cut -f2 -d' '`
    DATE=`date +%Y%m%d`
    docker build -f Dockerfile.ci --build-arg VERSION=$VERSION -t $DOCKER_REPO .
    docker tag $DOCKER_REPO $DOCKER_REPO:$DATE
    if [ "$TRAVIS_TAG" != "" ]; then
        docker tag $DOCKER_REPO:$TAG $DOCKER_REPO:$TRAVIS_TAG
    fi
    #docker tag $DOCKER_REPO $DOCKER_REPO:$TRAVIS_COMMIT
    #docker tag $DOCKER_REPO $DOCKER_REPO:travis-$TRAVIS_BUILD_NUMBER
    if [ "$TRAVIS_BRANCH" = "master" ] && [ "$TRAVIS_PULL_REQUEST" = "false" ]; then
        docker login -u "$DOCKER_USER" -p "$DOCKER_PASS"
        docker push $DOCKER_REPO
    fi
# windows, linux, osx pre-deploy
elif [ "$TARGET" != "docs" ]; then
    mkdir upload
    zip -r upload/pg_tileserv_${TAG}_${TARGET}.zip ${BINARY} README.md LICENSE.md assets/ config/
fi
