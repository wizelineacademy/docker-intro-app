#!/usr/bin/env bash

echo "******************"
docker-compose -f Backend/docker-compose.05.yml -f Frontend/docker-compose.05.yml up
