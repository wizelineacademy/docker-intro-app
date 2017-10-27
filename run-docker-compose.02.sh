#!/usr/bin/env bash

echo "******************"
docker-compose -f  docker-compose-backend.02.yml -f  docker-compose-frontend.02.yml config
echo "******************"
docker-compose -f  docker-compose-backend.02.yml -f docker-compose-frontend.02.yml up
