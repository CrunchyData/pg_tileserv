#!/bin/bash

echo "docs deploy ran"

if [ "$TARGET" = "docs" ]; then
    echo "in docs deploy branch: $TRAVIS_BRANCH"

    git add -f docs
    git status
    git config user.name "Travis CI"
    git config user.email "travis@travis-ci.org"
    git commit --message "Auto deploy from Travis CI"
    git remote add deploy "https://$GITHUB_TOKEN@github.com/pramsey/pg_tileserv"
    git push deploy docbuild
fi

