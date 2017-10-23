#!/usr/bin/env bash

echo "******************"
docker-compose -f Backend/docker-compose.09.yml -f Frontend/docker-compose.09.yml up
