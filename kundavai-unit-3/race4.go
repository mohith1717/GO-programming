package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup, name string) {
	fmt.Println("Hello,", name)
	wg.Done() // Signal that the goroutine is finished
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Add 2 goroutines to the wait group

	go sayHello(&wg, "Alice")
	go sayHello(&wg, "Bob")

	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("All greetings sent.")
}
