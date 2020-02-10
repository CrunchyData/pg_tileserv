#!/bin/bash

echo "docs deploy ran"

if [ "$GITHUB_TOKEN" != "" ]; then
    echo "Running documentation deploy script in $TRAVIS_BRANCH"
    git remote add deploy "https://$GITHUB_TOKEN@github.com/pramsey/pg_tileserv"
    git checkout -b docbuild
    git add -f docs
    git status
    git config user.name "Travis CI"
    git config user.email "travis@travis-ci.org"
    git commit --message "Auto deploy from Travis CI"
    git push deploy docbuild
fi

