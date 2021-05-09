package main

import (
	"fmt"
	"time"
)

func printjob(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%s", s)
	}
}
func main() {

	fmt.Println("[START]")
	go printjob("A")
	go printjob("B")
	go printjob("C")
	time.Sleep(1 + time.Second)

	fmt.Println("[Finish]")
}
