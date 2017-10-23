#!/usr/bin/env bash

echo "******************"
docker-compose -f Backend/docker-compose.03.yml -f Frontend/docker-compose.03.yml config
echo "******************"
docker-compose -f Backend/docker-compose.03.yml -f Frontend/docker-compose.03.yml up

