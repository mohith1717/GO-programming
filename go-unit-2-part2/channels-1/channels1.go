package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[3:6], c)
	go sum(s[0:3], c)
	x := <-c // receive from c
	y := <-c // receive from c

	fmt.Println(x, y, x+y)

	close(c)
}
