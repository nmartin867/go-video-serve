BIN=     ./bin/
SRC?=    $(shell find . -type f -name '*.go' -not -path "./vendor/*")
TARGET=  app-server

default: help

help:	## show this help
	@grep -E '^(.+):\s.*##\s(.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

all:	## clean, format, build and test
	make clean-all
	make gofmt
	make build
	make test

build:	## build all platforms
	GOARCH=amd64 GOOS=darwin go build -o ${BIN}${TARGET} main.go

run:   ## run target
	 go run main.go

clean:
	go clean
	rm -rf ${BIN}

clean-all:	## remove all generated artifacts and clean all build artifacts
	go clean -i ./...
	make clean

tidy:
	go mod tidy -v

fmt:    ## format the go source files
	go fmt ./...
	make tidy
	staticcheck ./...

