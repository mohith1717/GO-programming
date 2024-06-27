package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
	// Race condition introduced: directly access counter without mutex
	counter++
	fmt.Println("Counter:", counter)
	wg.Done() // Signal that the goroutine is finished
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10) // Add 10 goroutines to the wait group

	for i := 0; i < 10; i++ {
		go increment(&wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final Counter:", counter)
}
