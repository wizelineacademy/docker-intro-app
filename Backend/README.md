# docker-intro-backend
Sample API in golang for our docker intro session

### Installation

Install Dependencies 
``` bash
$ make dep
```


### Run Application

Run The api
``` bash
$ make run
```

### Modify DB connection

The app has the following config env variables

MONGODB_PORT_27017_TCP_ADDR : Address to the Mongo DB host
MONGODB_PORT_27017_TCP_PORT : Port to connect to the Mongo DB


### Test Application

test your applicaiton
``` bash
$ make test
```

###
check Docker branch for details on how to run with docker image

### Pending Taks

- [x] Create CRUD for TODOs
- [ ] Unit Testing
- [ ] Handle Environemts via ENV variables
- [ ] Create Docker File