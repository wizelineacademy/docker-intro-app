#!/usr/bin/env bash

docker version
echo "******************"

docker-compose version
echo "******************"


docker-compose -f docker-compose.10.yml config
echo "******************"

docker ps -a | awk '{ print $1,$2 }' | grep academy | awk '{print $1 }' | xargs -I {} docker stop {}
docker ps -a | awk '{ print $1,$2 }' | grep academy | awk '{print $1 }' | xargs -I {} docker rm {}

docker-compose  -f docker-compose.10.yml up -d
echo "******************"

docker ps
echo "******************"
