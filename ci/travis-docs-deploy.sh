#!/bin/bash

echo "docs deploy ran"

if [ "$TARGET" = "docs" ]; then
    echo "in docs deploy branch"
    git add -f docs
    git status
fi

