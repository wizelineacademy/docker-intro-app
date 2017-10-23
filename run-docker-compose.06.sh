#!/usr/bin/env bash

echo "******************"
docker-compose -f Backend/docker-compose.06.yml -f Frontend/docker-compose.06.yml up
