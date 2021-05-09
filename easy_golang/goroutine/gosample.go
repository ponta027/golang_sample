package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printjob(s string) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%s", s)
	}
}

func main() {

	wg.Add(3)
	go printjob("A")
	go printjob("B")
	go printjob("C")
	wg.Wait()
}
