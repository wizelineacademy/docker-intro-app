#!/usr/bin/env bash

docker version
echo "******************"

docker-compose version
echo "******************"

docker-compose -f docker-compose.10.yml config
echo "******************"

docker-compose -f docker-compose.10.yml up -d
echo "******************"

docker ps
echo "******************"
