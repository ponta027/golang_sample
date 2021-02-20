NAME :=gowiki 
VERSION := $(gobump show -r )
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"




bin/%: cmd/%/main.go
	go build -ldflags $(LDFLAGS) -o $@ $<

build : bin/gowiki



clean:
	rm bin/*
