package main

import (
	"flag"
	"ponta027.dip.jp/rest"
)


func main() {
	var (
		configPath   = flag.String("path", "config.json", "specify config path")
	)
	flag.Parse()
	rest.Request(*configPath)
}
