package main

import (
	"fmt"
	"os"
)

func main() {

	fname := "sample.dat"
	var file *os.File
	var err error

	file, err = os.Create(fname)
	if err != nil {
		_ = fmt.Errorf("%s is not open ", fname)
	}
	txt := "this is sample"
	file.Write(([]byte)(txt))
	file.Close()

	const BUFSIZE = 1024
	file, err = os.Open(fname)
	if err != nil {
		_ = fmt.Errorf("%s is not open ", fname)
	}
	defer file.Close()
	buf := make([]byte, BUFSIZE)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			_ = fmt.Errorf("%s if read error ", fname)
			break
		}
		fmt.Printf(string(buf[:n]))
	}
}

//
