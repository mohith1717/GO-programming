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

	// Simple go synchronous program
	// without any concurrency
	execute("First")

	// when this func is called it executes
	// and then passes on to next line
	execute("Second")

	// after the first second is to be executed
	// the problem is, execute function will
	// never finish execution, its trapped
	// in an infinite loop so the program will
	// never move to second execution.
	fmt.Println("program ends successfully")

	// if I'm wrong and both first and second
	// execute, then this line should print too
	// check the output
}
