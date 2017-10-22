GOPATH := ${CURDIR}
PROJECT_ROOT := $(shell pwd)
PROJECT_NAME := api
BIN_NAME := $(PROJECT_NAME)

clean:
	@rm -rf bin pkg package

init: clean
	@mkdir bin

dep:
	go get github.com/Pepegasca/goop
	@goop update

build:
	@goop exec go build -o bin/$(BIN_NAME) 

run: build
	bin/$(BIN_NAME)

fmt:
	@gofmt -s -l -w src

l: fmt run

test:
	@goop exec go test -v  -timeout 30s  ./...
