#!/usr/bin/env bash

docker rmi $(docker images -q)
docker rm $(docker ps -a -q)