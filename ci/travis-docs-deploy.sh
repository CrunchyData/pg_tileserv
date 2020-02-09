#!/bin/bash

echo "docs deploy ran"

if [ "$TARGET" = "docs" ]; then
    echo "in docs deploy branch"
    git add -f docs
    git status
    git commit -m 'travis doc build'
    echo branch: $TRAVIS_BRANCH
    git push docbuild --force
fi

