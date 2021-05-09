package main

import (
	"fmt"
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
	ifile, ierr := os.OpenFile(fname, os.O_RDONLY, 0444)
	if ierr != nil {
		log.Fatal(ierr)

	}

	var v int
	var str string
	for {

		n, err := fmt.Fscanf(ifile, "%d %s", &v, &str)
		if n == 0 || err != nil {

			break
		} else {
			fmt.Printf("%03d %s\n", v, str)
		}

	}

}
