#!/usr/bin/env bash

echo "******************"
docker-compose -f Backend/docker-compose.04.yml -f Frontend/docker-compose.04.yml up
