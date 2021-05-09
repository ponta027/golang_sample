package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("引数を２個指定してください")
		os.Exit(1)
	}

	src, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}
