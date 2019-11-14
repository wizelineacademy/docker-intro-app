@title[3-tier-app]

### How does our containers communicate?

---

@snap[north]
#### Let's create a Docker network
@snapend

```sh
# create a network
$ docker network create app

# list networks
$ docker network ls
```

---

### Running our database

---

@snap[north]
#### Run mongo:4
@snapend

```sh
$ docker run \
  --detach \
  --name mongo \
  --network app \
  mongo:4
```

@snap[south text-07]
Note how we do not publish the network port
@snapend

---

### Building and running our backend

---

@snap[north]
#### Let's create a Dockerfile
@snapend

```sh
# move to the Backend/ directory
$ cd Backend/

# open the Dockerfile
$ open Dockerfile
```

---

---?code=slides/3-tier-app/backend/Dockerfile&lang=dockerfile
@snap[south span-100]
@[1-11](copy and paste the instructions in your Dockerfile)
@snapend

---

@snap[north]
#### Build and run api:multi-stage
@snapend

```sh
$ docker build --tag api:multi-stage .

$ docker run \
  --detach \
  --name api \
  --network app \
  --publish 8000:8000 \
  --rm \
  api:multi-stage
```

---

### Running our frontend

---

@snap[north]
#### Run web:multi-stage
@snapend

```sh
$ docker run \
  --detach \
  --name web \
  --network app \
  --publish 80:80 \
  --rm \
  web:multi-stage
```

---

### Try it out!
@fa[desktop](localhost:80)

---

### Clean up
```sh
$ docker rmi \
  web:simple \
  web:cache \
  web:multi-stage \
  app:multi-stage \
  mongo:4

$ docker network rm app
```
