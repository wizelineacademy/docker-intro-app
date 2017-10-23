GOPATH =$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
PROJECT_ROOT := $(shell pwd)
PROJECT_NAME := api
BIN_NAME := $(PROJECT_NAME)

clean:
	@rm -rf bin pkg package vendor .vendor

init: clean
	@mkdir bin

dep:
	@GOBIN=$(GOBIN) go get github.com/Pepegasca/goop
	./bin/goop update

build:
	@./bin/goop exec go build -o bin/$(BIN_NAME) 

run: build
	bin/$(BIN_NAME)

fmt:
	@gofmt -s -l -w src

l: fmt run

test:
	@./bin/goop  exec go test -v  -timeout 30s  ./...

