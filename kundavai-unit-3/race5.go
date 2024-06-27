package main

import (
	"fmt"
	"sync"
)

var counter int

func sayHello(wg *sync.WaitGroup, name string) {
	counter++ // Accessing shared variable without protection
	fmt.Println("Hello,", name, " - Counter:", counter)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Add 2 goroutines to the wait group

	go sayHello(&wg, "Alice")
	go sayHello(&wg, "Bob")

	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("All greetings sent.")
}
