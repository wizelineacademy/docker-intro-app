#!/usr/bin/env bash

docker version
echo "******************"

docker-compose version
echo "******************"


docker-compose -f docker-compose.yml config
echo "******************"

docker ps -a | awk '{ print $1,$2 }' | grep academy | awk '{print $1 }' | xargs -I {} docker stop {}
docker ps -a | awk '{ print $1,$2 }' | grep academy | awk '{print $1 }' | xargs -I {} docker rm {}

docker-compose  -f docker-compose.yml up -d
echo "******************"

docker ps
echo "******************"
