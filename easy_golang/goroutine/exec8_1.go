package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			fmt.Printf("+")
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			fmt.Printf("-")
			time.Sleep(time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Printf("Finish")
}
