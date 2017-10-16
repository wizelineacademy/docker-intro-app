# docker-intro-backend
Sample API in golang for our docker intro session

### Installation
install Goop via 

go get github.com/Pepegasca/goop
then

Install Dependencies 
``` bash
$ make dep
```


### Run Application

Run The api
``` bash
$ make run
```


### Test Application

test your applicaiton
``` bash
$ make test
```

### Pending Taks

- [x] Create CRUD for TODOs
- [ ] Unit Testing
- [x] Handle Environemts via ENV variables
- [x] Create Docker File

### Docker build run

Have Mongo DB container running

``` bash
docker run -d -p 0.0.0.0:27017:27017 --name sample-mongo mongo
```
Build your image from the Dockerfile

``` bash
docker build -t backend_app .
```

Run your App image linked to mongodb

``` bash
docker run -d --name backend -p 0.0.0.0:8000:8000 --link sample-mongo:mongodb backend_app
```