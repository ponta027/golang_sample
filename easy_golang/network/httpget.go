package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	fmt.Println("[Start]")

	if len(os.Args) < 2 {
		fmt.Println(" lack argument")
		fmt.Println("[END]")
		os.Exit(1)
	}

	url := os.Args[1]

	if !strings.HasPrefix(url, "http") {
		fmt.Println("[END]")
		os.Exit(1)
	}

	fmt.Println(" get %s", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("[END]")
		os.Exit(1)
	}

	src, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Println("[END]")
		os.Exit(1)
	}
	fmt.Println("%s", string(src))
	fmt.Println("[END]")

}
