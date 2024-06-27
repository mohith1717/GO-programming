package main

import "fmt"

func main() {

	ss := make(chan string, 2)
	ss <- "Hello"
	ss <- "World"

	//   ss <- "Blocked"

	fmt.Println(<-ss)
	fmt.Println(<-ss)
}
