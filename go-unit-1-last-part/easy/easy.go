package main

import "fmt"

func main() {

	defer fmt.Println("Morning")
	defer fmt.Println("Good")

	some(3)
	fmt.Println("hello")
}

func some(x int) int {

	defer fmt.Println("1st")
	defer fmt.Println("2nd")
	fmt.Println(x)
	defer fmt.Println("3rd")
	return 45
}
