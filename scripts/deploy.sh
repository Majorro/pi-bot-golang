#!/bin/bash

git fetch

UPSTREAM='origin/master'
LOCAL=$(git rev-parse @)
REMOTE=$(git rev-parse "$UPSTREAM")
BASE=$(git merge-base @ "$UPSTREAM")

if [ $LOCAL = $REMOTE ]; then
    echo "Up-to-date"
elif [ $LOCAL = $BASE ]; then
    echo "Need to pull"
    docker-volume-snapshot create pi-bot_db-data db-backup.tar.gz
    git pull
    docker compose up -d --no-deps --build pibot
else
    echo "???"
fi
