#!/usr/bin/env bash

docker version
echo "******************"

docker-compose version
echo "******************"

docker-compose -f docker-compose.05.yml config
echo "******************"

docker-compose -f docker-compose.05.yml up -d
echo "******************"

docker ps
echo "******************"
