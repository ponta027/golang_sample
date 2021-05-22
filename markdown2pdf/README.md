# README

## 動作確認環境

```
Linux raspberrypi3 4.14.98-v7+ #1200 SMP Tue Feb 12 20:27:48 GMT 2019 armv7l GNU/Linux
```

## setup   

* install wkhtmltopdf

See . https://wkhtmltopdf.org/downloads.html

* install pdftk

> apt-get install pdftk


## build configuration

> go mod edit -replace markdown2pdf/md2pdf@v1.0.0=./md2pdf


## how to build

* default environment

> make build 

* windows environment

> make build GOOS="windows" GOARCH=amd64


## command

```
Usage of ./bin/md2pdf:
  -i string
    	input markdown file (default "test.md")
  -o string
    	output pdf file (default "test.pdf")
  -p string
    	password.if password is empty, not protect
```

## make configuration

```
build:             build binary 
                   ex. cross compile  GOOS=linux GOARCH=arm64 make build
clean:             clean
devel-deps:        download dependent module 
help:              show help
init:              initialize project cconfiguration
```


