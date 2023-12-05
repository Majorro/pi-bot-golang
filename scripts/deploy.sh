#!/bin/bash

git fetch

UPSTREAM='origin master'
LOCAL=$(git rev-parse @)
REMOTE=$(git rev-parse "$UPSTREAM")
BASE=$(git merge-base @ "$UPSTREAM")

if [ $LOCAL = $REMOTE ]; then
    echo "Up-to-date"
elif [ $LOCAL = $BASE ]; then
    echo "Need to pull"
    git pull $UPSTREAM
    chmod +x ./scripts/deploy.sh
    docker compose up -d --no-deps --build pibot
else
    echo "???"
fi
