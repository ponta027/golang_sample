package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fname := "./sample.dat"
	txt := "this is sample"
	werr := ioutil.WriteFile(fname, ([]byte)(txt), 0666)
	if werr != nil {
		_ = fmt.Errorf("%s is not writable ", fname)
	}
	data, rerr := ioutil.ReadFile(fname)
	if rerr != nil {
		_ = fmt.Errorf("%s is not readable ", fname)
	} else {
		fmt.Print(string(data))
	}

}
