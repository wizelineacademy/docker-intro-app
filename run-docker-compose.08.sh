#!/usr/bin/env bash

echo "******************"
docker-compose -f Backend/docker-compose.08.yml -f Frontend/docker-compose.08.yml up
