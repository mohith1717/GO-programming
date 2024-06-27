package main

import (
	"fmt"
	"time"
)

func execute(some string) {

	// initializing a infinite for loop
	for i := 1; true; i++ {

		// prints string and number
		fmt.Println(some, i)

		// this makes the program sleep for
		// 100 milliseconds, wiz 10 seconds
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	// Simple go program with concurrency
	go execute("First")

	// Placing the go command in front of the
	// func call simply creates a goroutine
	go execute("Second")

	// The second goroutine, you may think that the
	// program will now run with lightning speed
	// But, both goroutines go to the background
	// and result in no output at all Because the
	// program exits after the main goroutine
	fmt.Println("Program ends successfully")

	// This statement will now be executed
	// and nothing else will be executed
	// check the output
}
