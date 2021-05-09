package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("+")
		}
		done <- true
	}()
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("/")
	}

	<-done
	fmt.Printf("Finish")
}
