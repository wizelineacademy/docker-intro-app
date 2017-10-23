#!/usr/bin/env bash

echo "******************"
docker-compose -f Backend/docker-compose.01.yml -f Frontend/docker-compose.01.yml config

