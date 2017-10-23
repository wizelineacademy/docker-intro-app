#!/usr/bin/env bash

docker stop $(docker ps -aq)
docker rmi $(docker images -q)
docker rm $(docker ps -a -q)