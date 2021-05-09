package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fname := "./sample2.dat"
	var err error
	var file *os.File
	file, err = os.Create(fname)
	if err != nil {
		_ = fmt.Errorf("%s is not open", fname)
	}
	var txt string
	for i := 0; i < 5; i++ {

		txt = "this is sample " + strconv.Itoa(i+1) + " line\n"
		file.Write(([]byte)(txt))
	}

	file.Close()

	file, err = os.Open(fname)
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			_ = fmt.Errorf("%s is not read ", fname)
			break
		}
		fmt.Println(sc.Text())
	}

}
