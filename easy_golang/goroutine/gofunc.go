package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for _, c := range "ABC" {
		ch := c
		go func() {
			//            wg.Add(1)
			for i := 0; i < 10; i++ {
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("%c", ch)
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Finish")

}
