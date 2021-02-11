NAME := rest 
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

bin/%: cmd/%/main.go 
	go build -ldflags $(LDFLAGS) -o $@ $<

## build binary 
## ex. cross compile  GOOS=linux GOARCH=arm64 make build
.PHONY: build
build: bin/rest_sample

test: 
	go test -v ponta027.dip.jp/rest -coverprofile=artifact/cover.out	¥
	go tool cover -html=artifact/cover.out -o artifact/cover.html


## show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)


## clean
.PHONY: clean
clean:
	rm bin/*
