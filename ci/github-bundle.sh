#!/bin/bash

# Exit on failure
set -e

echo "GITHUB_REF_NAME = $GITHUB_REF_NAME"
echo "MATRIX_OS = $MATRIX_OS"

if [ "$MATRIX_OS" = "ubuntu-latest" ]; then
    TARGET="linux"
elif [ "$MATRIX_OS" = "macos-latest" ]; then
    TARGET="macos"
elif [ "$MATRIX_OS" = "windows-latest" ]; then
    TARGET="windows"
else
    echo "ERROR: Unsupported OS, $MATRIX_OS"
    exit 1
fi

if [ "$GITHUB_REF_NAME" = "master" ]; then
    TAG="latest"
else
    TAG=$GITHUB_REF_NAME
fi

if [ "$MATRIX_OS" = "windows-latest" ]; then
    BINARY=pg_tileserv.exe
else
    BINARY=pg_tileserv
fi

PAYLOAD="${BINARY} README.md LICENSE.md assets/ config/"
ZIPFILE="pg_tileserv_${TAG}_${TARGET}.zip"

echo "ZIPFILE = $ZIPFILE"
echo "PAYLOAD = $PAYLOAD"

mkdir upload
#zip -r upload/$ZIPFILE $PAYLOAD
7z a upload/$ZIPFILE $PAYLOAD
