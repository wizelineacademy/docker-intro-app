#!/usr/bin/env bash

echo "******************"
docker-compose -f docker-compose-backend.01.yml -f docker-compose-frontend.01.yml config

