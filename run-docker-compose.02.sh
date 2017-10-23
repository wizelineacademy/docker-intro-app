#!/usr/bin/env bash

echo "******************"
docker-compose -f Backend/docker-compose.02.yml -f Frontend/docker-compose.02.yml config
echo "******************"
docker-compose -f Backend/docker-compose.02.yml -f Frontend/docker-compose.02.yml up
