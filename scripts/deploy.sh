#!/bin/bash

# TODO: maybe use https://github.com/mag37/dockcheck with dockerfile build in cd pipeline

git fetch

UPSTREAM='origin/master'
LOCAL=$(git rev-parse @)
REMOTE=$(git rev-parse "$UPSTREAM")
BASE=$(git merge-base @ "$UPSTREAM")

if [ $LOCAL = $REMOTE ]; then
    echo "Up-to-date"
elif [ $LOCAL = $BASE ]; then
    echo "Need to pull"
    git pull
    docker-volume-snapshot create pi-bot_db-data db-backup.tar.gz &&
    docker compose up -d --no-deps --build pibot
else
    echo "???"
fi
