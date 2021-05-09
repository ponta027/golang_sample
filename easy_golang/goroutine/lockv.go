package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var a int
var mu sync.Mutex

func jobplus() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		t := a
		a = a + 1
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("plus %d->%d¥n", t, a)
		mu.Unlock()
	}
}
func jobminus() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		t := a
		a = a - 1
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("minus %d->%d¥n", t, a)
		mu.Unlock()
	}
}

func main() {
	a = 10
	wg.Add(2)
	go jobplus()
	go jobminus()
	wg.Wait()
}
