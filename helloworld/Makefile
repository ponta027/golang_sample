NAME := helloworld
VERSION := $(gobump show -r )
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"


## install dependencies
.PHONY: deps
deps:
	go get -v -d 

## setup
.PHONY: deps
devel-deps: deps
	go get github.com/jessevdk/go-flags

# build binary
bin/%: cmd/%/main.go 
	go build -ldflags $(LDFLAGS) -o $@ $<

## build binary
.PHONY: build
build: bin/helloworld

## show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)


## clean
.PHONY: clean
clean:
	rm bin/helloworld
