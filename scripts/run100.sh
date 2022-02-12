#!/bin/sh
for i in {1..500}
do
    battlesnake play -W 11 -H 11 --name wrappedhack --url http://0.0.0.0:7070 --name anti-blobby-live --url http://0.0.0.0:6969 -g wrapped
done
