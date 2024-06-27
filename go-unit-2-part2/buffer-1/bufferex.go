package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ss := make(chan string, 2)
	ss <- "Hello"
	ss <- "World"
	fmt.Println(<-ss)
	fmt.Println(<-ss)
}
