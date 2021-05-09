package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fname := "sample.dat"
	words := [...]string{"Dog", "Cat", "Pig", "Deer", "Racoons"}

	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	//
	for i, w := range words {
		fmt.Fprintf(file, "%03d %s\n", i+1, w)
	}
	file.Close()

	//
	txt, rerr := ioutil.ReadFile(fname)
	if rerr != nil {
		log.Fatal(rerr)
	}
	fmt.Printf(string(txt))
}
