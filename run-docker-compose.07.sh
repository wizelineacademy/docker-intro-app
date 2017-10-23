#!/usr/bin/env bash

echo "******************"
docker-compose -f Backend/docker-compose.07.yml -f Frontend/docker-compose.07.yml up
