package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	pool := semaphore.NewWeighted(2)
	go swim("A", pool)
	go swim("B", pool)
	go swim("C", pool)
	go swim("D", pool)

	time.Sleep(5 * time.Second) // For brevity, better use sync.WaitGroup
	log.Println("Main: Done, shutting down")
}

func swim(name string, pool *semaphore.Weighted) {
	log.Printf("%v: I want to swim\n", name)

	// In real applications, pass in your context such as HTTP request context
	ctx := context.Background()

	// We can also Acquire/Release more than 1
	// when the workloads consume different amount of resources
	if err := pool.Acquire(ctx, 1); err != nil {
		log.Printf("%v: Ops, something went wrong! I cannot acquire a lane\n", name)
		return
	}

	log.Printf("%v: I got a lane, I'm swimming\n", name)
	time.Sleep(time.Second)
	log.Printf("%v: I'm done. Releasing my lane\n", name)
	pool.Release(1)
}
