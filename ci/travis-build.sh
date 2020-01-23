#!/bin/bash

if [ "$TARGET" = "windows" ]; then
    env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v
else
    go build -v
fi


