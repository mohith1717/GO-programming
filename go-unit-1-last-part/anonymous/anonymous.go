package main

import "fmt"

func main() {

	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}

func substract() func() int {
	var y int
	return func() int {
		y++
		return y * y
	}
}

// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}


func somesquare() func() int{
	var x int
	return func()int{
		x++
		return x*x
	}
}