package main

import (
	"fmt"
	"time"
)

func main() {
	go MyFunction()
	// MyFunction()
	MyFunction()

	// Adding a delay to see the asynchronous behavior
	time.Sleep(6 * time.Second)
}

func MyFunction() int {
	fmt.Println("Entering MyFunction")
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting MyFunction")
	return 0
}
