#! /bin/bash

# set -x

# To monitor container locally you may use a tool such as:
# pip install glances[docker]

RUN_SECTION=$1

p_instructions(){
  COLOUR='\033[1;36m' # Light Cyan
  NC='\033[0m' # No Color

  echo -e "${COLOUR}\n---> $1\n${NC}"
}

get_container_id(){
  C_ID=$1
  if [ -z "$C_ID" ]; then exit; fi
  docker ps -a | grep $C_ID | awk '{ print $1 }'
}

if [ -z "$RUN_SECTION" ]; then
cat <<EOF

  usage: ./troubleshoot <section number>

  SECTIONS:
    1: build and run detached, then print logs
    2: run again but detach and then attach
    3: Override CMD and enter container
    4: Do: inspect,  stop,  rm,  images,  rmi

   NOTE: if one fails run them in order: 1, 2, 3..
EOF
  exit
fi

if [[ $RUN_SECTION == "1" ]]; then
  # Stop the container if it is already running in detached
  CONTAINER_ID=$(get_container_id docker-academy-backend)
  docker stop $CONTAINER_ID
  docker rm $CONTAINER_ID

  # Build and run the container in detached mode
  docker build -t docker-academy-backend -f Dockerfile.build .
  p_instructions "Finished build. Running container"
  docker run -it --name docker-academy-backend -p 8000:8000 -d docker-academy-backend
  p_instructions "Container is now running, waiting 30 seconds to proceed"

  # wait for the container to fail
  sleep 30

  # Get the container ID
  CONTAINER_ID=$(get_container_id docker-academy-backend)

  # Print the logs to get insights of errors.
  p_instructions "Print the logs to get insights of errors."
  set -x
  docker logs $CONTAINER_ID
  echo ""
  docker ps -a
  p_instructions "Contemplate above errors..."
  exit
fi

if [[ $RUN_SECTION == "2" ]]; then
  set -x
  # Run again but detach and then attach:
  p_instructions "Run again but detach and then attach:"
  CONTAINER_ID=$(get_container_id docker-academy-backend)
  docker stop $CONTAINER_ID
  docker rm $CONTAINER_ID
  docker run -it --name docker-academy-backend -p 8000:8000 -d docker-academy-backend
  CONTAINER_ID=$(get_container_id docker-academy-backend)
  p_instructions "Attaching to already running container:"
  docker attach $CONTAINER_ID
  echo ""
  docker ps -a
  p_instructions "Contemplate above errors..."
  exit
fi

# echo -e "\n---> Entering the container\n"
# docker exec -it $CONTAINER_ID /bin/bash

if [[ $RUN_SECTION == "3" ]]; then
  # Override CMD and enter container.
  p_instructions "Override CMD and enter container. we can run bin/api inside"
  set -x
  docker run -it --rm docker-academy-backend /bin/bash
  exit
fi

if [[ $RUN_SECTION == "4" ]]; then
  set -x
  p_instructions "Remove Containers and Images"
  IMAGE_NAME=docker-academy-backend
  CONTAINER_ID=$(get_container_id $IMAGE_NAME)
  docker inspect $CONTAINER_ID
  docker stop $CONTAINER_ID
  docker rm $CONTAINER_ID
  # Docker prune <object>
  docker images | head
  docker rmi $IMAGE_NAME
  docker ps -a
  exit
fi
