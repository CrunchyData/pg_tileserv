#!/bin/bash

# Exit on failure
set -e

DATE=`date +%Y%m%d`

echo "GITHUB_REF_NAME = $GITHUB_REF_NAME"
echo "GITHUB_HEAD_REF = $GITHUB_HEAD_REF"
echo "DOCKER_REPO = $DOCKER_REPO"
echo "DATE = $DATE"

if [ "$GITHUB_REF_NAME" = "master" ]; then
    TAG="latest"
else
    TAG=$GITHUB_REF_NAME
fi

if [ "$GITHUB_REF_NAME" = "master" ] && [ "$GITHUB_HEAD_REF"x = "x" ]; then
    echo "Logging in..."
    echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
    docker push $DOCKER_REPO
fi
