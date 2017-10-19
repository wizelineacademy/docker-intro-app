# todo-list

> docker intro sample front end app

## Build Setup

``` bash
# install dependencies
npm install

# serve with hot reload at localhost:9090
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report

# run unit tests
npm run unit

# run e2e tests
npm run e2e

# run all tests
npm test
```

For detailed explanation on how things work, checkout the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).

### Pending Taks

- [x] Create TODO list app
- [x] Update to use API instead of just state
- [ ] Unit Testing
- [x] Handle Environemts via ENV variables
- [x] Create Docker File

### Dockerized App

*You need to have your backend docker running by now*

Run the below command to build your Image
``` bash
docker build -t frontend .
```

Run the below command to Run your Image and Link it to your backend API
``` bash
docker run -d -p 0.0.0.0:80:80 --link backend:api --name frontend frontend
```