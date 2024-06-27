package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	// Creating a semaphore with a limit of 2
	pool := semaphore.NewWeighted(2)

	// We'll simulate four tasks trying to acquire resources concurrently
	for i := 0; i < 4; i++ {
		go swim(i, pool)
	}

	// Sleep for some time to allow goroutines to execute
	time.Sleep(5 * time.Second)

	log.Println("Main: Done, shutting down")
}

func swim(id int, pool *semaphore.Weighted) {
	log.Printf("Task %d: Trying to acquire a resource\n", id)

	// In real applications, pass in your context such as HTTP request context
	ctx := context.Background()

	// Acquire a resource from the pool
	if err := pool.Acquire(ctx, 1); err != nil {
		log.Printf("Task %d: Unable to acquire resource: %v\n", id, err)
		return
	}

	log.Printf("Task %d: Acquired a resource, performing task\n", id)

	// Simulating some task
	time.Sleep(time.Second)

	log.Printf("Task %d: Task completed, releasing resource\n", id)

	// Release the acquired resource
	pool.Release(1)
}
