#!/bin/bash

if [ "$TARGET" = "windows" ]; then
	sudo apt install gcc-mingw-w64 libc6-dev-i386
fi

if [ "$TARGET" = "docs" ]; then
    VER=0.64.1
    FILE=hugo_${VER}_Linux-64bit.deb
	URL=https://github.com/gohugoio/hugo/releases/download/v${VER}/${FILE}
    curl -LO $URL
    sudo dpkg -i ${FILE}
fi

