#!/bin/bash

docker run -it --rm --name todo-backend -p 8000:8000 --link sample-mongo:mongodb todo-backend
