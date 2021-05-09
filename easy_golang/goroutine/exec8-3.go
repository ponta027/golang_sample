package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	str := make(chan string)

	go func() {
		for {
			s, _ := <-str
			fmt.Printf("%s", strings.ToUpper(s))
			if s == "quit" {
				break

			}
		}
	}()

	for {
		fmt.Printf("Input:")
		fmt.Scanf("%s", &s)
		if s == "" {
			continue
		}
		str <- s
		if s == "quit" {
			break
		}

	}
	fmt.Printf("Finish")
}
