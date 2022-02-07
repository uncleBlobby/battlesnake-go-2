#!/bin/sh
for i in {1..500}
do
    battlesnake play -W 11 -H 11 --name 20blobby2 --url http://0.0.0.0:8080 --name anti-blobby --url http://0.0.0.0:6969 -g royale
done
