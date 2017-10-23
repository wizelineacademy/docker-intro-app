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


# PATH=/Users/charlycastro/Library/Python/3.6/bin:/usr/local/sbin:/opt/local/bin:/opt/local/sbin:/opt/local/bin:/opt/local/sbin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/MacGPG2/bin:/usr/local/go/bin:/Users/charlycastro/go_work/bin:/usr/local/mysql/bin/:/usr/local/go/bin:/Users/charlycastro/go_work/bin:/usr/local/mysql/bin/