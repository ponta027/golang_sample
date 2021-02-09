package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		i = flag.Int("int", 0, "int flag")
	)
	flag.Parse()
	fmt.Println(*i)
	fmt.Println("Hello World")
}
