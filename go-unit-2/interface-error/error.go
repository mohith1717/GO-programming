package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("At %v,\n %s", e.When, e.What)
	// return fmt.Println("Hello brother")
}

func run() error {
	return &MyError{
		time.Now(),
		"\nit didn't work",
	}
}

func main() {

	i, err := strconv.Atoi("42A")
	// if err != nil {
	// 	fmt.Printf("couldn't convert number: %v\n", err)
	// 	return
	// }
	fmt.Println("Converted integer:\n", i)

	if err = run(); err != nil {
		fmt.Println(err)
	}

}
