package main

import (
	"fmt"
	"time"
)

func my_function(x int, y string) {

	for i := 1; i <= x; i++ {

		fmt.Println(y)

		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	// Simple go routines with concurrency
	go my_function(20, "first")

	go my_function(20, "second")

	time.Sleep(time.Millisecond * 1000)

	fmt.Println("program ends successfully")

}
