package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int32

	// Create a goroutine to increment the counter.
	go func() {
		for i := 0; i < 5; i++ {
			atomic.AddInt32(&counter, 1)
			fmt.Printf("Incremented: %d\n", atomic.LoadInt32(&counter))
			time.Sleep(time.Millisecond)
		}
	}()

	// Create a goroutine to decrement the counter.
	go func() {
		for i := 0; i < 5; i++ {
			atomic.AddInt32(&counter, -1)
			fmt.Printf("Decremented: %d\n", atomic.LoadInt32(&counter))
			time.Sleep(time.Millisecond)
		}
	}()

	// Wait for the goroutines to finish.
	time.Sleep(2 * time.Second)

	fmt.Printf("Final Value: %d\n", atomic.LoadInt32(&counter))
}
