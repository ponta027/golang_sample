NAME := webclient
VERSION := $(gobump show -r )
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"
GOOS := linux
GOARGCH := arm64

## build binary 
## ex. cross compile  GOOS=linux GOARCH=arm64 make build
.PHONY: build
build: bin/${NAME}



## initialize project cconfiguration
.PHONY: init
init: 
	mkdir -p cmd/${NAME}
	mkdir -p  ${NAME}
	go mod init ${NAME}


## download dependent module 
.PHONY: deps
devel-deps: deps
	go get github.com/russross/blackfriday


bin/%: cmd/%/main.go 
	go build -ldflags $(LDFLAGS) -o $@ $<

test: 
	go test -v webclient/webclient
#	go test -v webclient -coverprofile=artifact/cover.out	¥
#	go tool cover -html=artifact/cover.out -o artifact/cover.html


## show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)


## clean
.PHONY: clean
clean:
	rm bin/*
