#!/bin/bash

echo "docs deploy ran"

if [ "$GITHUB_TOKEN" != "" ]; then
    LOCAL_BRANCH=docs-$TRAVIS_BUILD_NUMBER
    echo "Running documentation deploy script in $TRAVIS_BRANCH"
    git remote add deploy "https://$GITHUB_TOKEN@github.com/CrunchyData/pg_tileserv.git"
    git checkout -b $LOCAL_BRANCH
    git branch -v
    git add -f docs
    git status
    git config user.name "Travis CI"
    git config user.email "travis@travis-ci.org"
    git commit --message "Auto deploy from Travis CI"
    git push deploy $LOCAL_BRANCH:master
fi

