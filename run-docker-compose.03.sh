#!/usr/bin/env bash

echo "******************"
docker-compose -f docker-compose-backend.03.yml -f docker-compose-frontend.03.yml config

echo "******************"
docker ps -a | awk '{ print $1,$2 }' | grep academy | awk '{print $1 }' | xargs -I {} docker stop {}
docker ps -a | awk '{ print $1,$2 }' | grep academy | awk '{print $1 }' | xargs -I {} docker rm {}

docker-compose  -f docker-compose-backend.03.yml -f  docker-compose-frontend.03.yml up

