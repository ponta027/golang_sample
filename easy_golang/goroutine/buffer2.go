package main

import (
	"fmt"
	"time"
)

func main() {
	// not change behavior when buffer size specified.
	done := make(chan int)
	//done := make(chan int,3)
	go func() {
		done <- 1
	}()
	go func() {
		done <- 2
	}()
	go func() {
		done <- 3
	}()

	time.Sleep(1 * time.Second)
	for i := 0; i < 3; i++ {
		fmt.Printf("done:%d", <-done)
	}
}
