package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/sync/semaphore"
)

func main() {
	ctx := context.TODO()

	// Create a weighted semaphore with 3 tokens
	sem := semaphore.NewWeighted(3)

	// Acquire tokens
	for i := 0; i < 5; i++ {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Fatalf("Failed to acquire semaphore: %v", err)
		}
		fmt.Printf("Acquired token %d\n", i+1)
	}

	// Release tokens
	sem.Release(3)

	fmt.Println("Released tokens")

	// Output:
	// Acquired token 1
	// Acquired token 2
	// Acquired token 3
	// Failed to acquire semaphore: context deadline exceeded
	// Failed to acquire semaphore: context deadline exceeded
	// Released tokens
}
