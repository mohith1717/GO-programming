package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	mutex.Lock()         // Acquire the lock before accessing the counter
	defer mutex.Unlock() // Release the lock after accessing the counter
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
