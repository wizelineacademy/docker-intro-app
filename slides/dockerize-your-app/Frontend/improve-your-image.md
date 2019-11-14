### Now let's build our image in the container
@title[Dockerize you App]
---?code=slides/dockerize-your-app/Frontend/Dockerfile.clean&title=Build%20Dockerfile
@snap[span-90]
@[1-2](Add our code to the nginx image)
@[4-9](This installs node 8)
@[10-12](This install dependencies, builds and copies to be served)
@snapend
