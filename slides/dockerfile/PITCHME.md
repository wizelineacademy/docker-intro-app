@title[Dockerfile]

## Creating a Docker image

---

@quote[Docker can build images automatically by reading the instructions from a Dockerfile](https://docs.docker.com/engine/reference/builder/)

---

@snap[north]
#### Let's create a Dockerfile
@snapend

```sh
# move to the Frontend/ directory
$ cd Frontend/

# open the Dockerfile
$ open Dockerfile
```

---?code=slides/dockerfile/simple/Dockerfile&lang=dockerfile
@snap[south span-100]
@[1-15](copy and paste the instructions in your Dockerfile)
@[1](defines a variable that users can pass at build-time)
@[2](sets the base image for subsequent instructions)
@[3](adds metadata to an image)
@[4](sets the user name and optionally the user group)
@[6](sets the working directory)
@[7-8](copies files and directories to the container's filesystem)
@[10](sets the environment variable for the build- and run-time)
@[11](executes any command and commits the result in the image)
@[13](informs Docker that the container listens on the specified network ports at runtime)
@[14](allows you to configure a container that will run as an executable)
@[14](only one ENTRYPOINT instruction can be added)
@[15](provides defaults for an executing container)
@[15](only one CMD instruction can be added)
@snapend

---

@snap[north]
#### Build and run web:simple
@snapend

```sh
# build an image from the Dockerfile
# docker build [--tag name:tag] <path>
$ docker build --tag web:simple .

# run a container
# docker run [--publish host:container] <name:tag>
$ docker run --publish 80:9090 web:simple
```

---

### How does Docker get the files?

---

@quote[...all recursive contents of files and directories in the current directory are sent to the Docker daemon as the build context](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)

@snap[south span-60]
```sh
# docker build [--tag name:tag] <path>
```
@snapend

---

@snap[north]
#### Let's create a dockerignore
@snapend

```sh
# open the dockerignore
$ open .dockerignore
```

---

@snap[north]
#### Ignoring files
@snapend

```docker
# dependencies
node_modules/

# build files
dist/

# Git files
.git/

# log files
*.log

# Docker files
Dockerfile
```

@snap[span-100 south])
@[1-14](copy and paste the instructions in your .dockerignore)
@[1-14](similar to .gitignore)
@snapend

---

@snap[north]
#### Docker layers
@snapend

- A Docker image is built up from a series of layers
- Each layer represents an instruction in the image’s Dockerfile
- Each layer except the very last one is read-only

---

---?code=slides/dockerfile/layers/Dockerfile&lang=dockerfile
@snap[north]
#### Example
@snapend

@snap[south span-100]
@[1-4](Only these instructions create layers)
@[1](FROM creates a layer from the ubuntu:18.04 Docker image)
@[2](COPY adds files from your Docker client’s current directory)
@[3](RUN builds your application with make)
@[4](CMD specifies what command to run within the container)
@snapend

---

### Reducing build time using cache

---

---?code=slides/dockerfile/cache/Dockerfile&lang=dockerfile
@snap[south span-100]
@[1-16](copy and paste the instructions in your Dockerfile)
@[7-12](first installs dependencies, and then copies files)
@snapend

---

@snap[north]
#### Build and run web:cache
@snapend

```sh
# build an image from the Dockerfile
# docker build [--tag name:tag] <path>
$ docker build --tag web:cache .

# run a container
# docker run [--publish host:container] <name:tag>
$ docker run --publish 80:9090 web:cache
```

---

### How big is our image?

---

@snap[north]
#### Image size
@snapend

@snap[west span-60]
```sh
# get local Docker images
$ docker images
```
@snapend

@snap[east span-40]
@box[bg-blue rounded text-07](web:cache # 401MB)
@snapend

@snap[south-east span-40]
@box[bg-blue rounded text-07](web:simple # 400MB)
@snapend

---

---?code=slides/dockerfile/multi-stage/Dockerfile&lang=dockerfile
@snap[north]
#### Reducing image size
@snapend

@snap[south span-100]
@[1-17](copy and paste the instructions in your Dockerfile)
@[1-14](first image builds the static content)
@[15-17](second image serves the static content)
@snapend

---

@snap[north]
#### Build and run web:multi-stage
@snapend

```sh
# build an image from the Dockerfile
# docker build [--tag name:tag] <path>
$ docker build --tag web:multi-stage .

# run a container
# docker run [--publish host:container] <name:tag>
$ docker run --publish 80:80 web:multi-stage
```

---

### Let's compare the images

---

@snap[north]
#### Image size
@snapend

@snap[west span-60]
```sh
# get local Docker images
$ docker images
```
@snapend

@snap[east span-40]
@box[bg-blue rounded text-07](web:multi-stage # 23.2MB)
@snapend

@snap[south-east span-40]
@box[bg-blue rounded text-07](web:cache # 401MB)
@snapend
