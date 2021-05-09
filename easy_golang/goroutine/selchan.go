package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("START")
	ch1 := make(chan rune)
	ch2 := make(chan int)
	done := make(chan bool)
	go func() {
		s := "ABCDEFG"
		for _, c := range s {
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("[ch1]%c¥r¥n", c)
			ch1 <- c
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 8; i++ {
			time.Sleep(8 * time.Millisecond)
			fmt.Printf("[ch2]%c¥r¥n", i)
			ch2 <- i
		}
		done <- true
	}()

	defer fmt.Printf("END")
	count := 0
	for {

		select {
		case r := <-ch1:
			fmt.Printf("[ch1][r]%c¥n", r)
		case n := <-ch2:
			fmt.Printf("[ch2][r]%c¥n", n)
		case <-done:
			count++
			if count > 1 {
				return
			}

		}
	}
}
