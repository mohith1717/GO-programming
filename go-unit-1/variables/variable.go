package main

import "fmt"

func main() {

	var i int
	var i1 int = 32

	var s string
	var s1 string = "Hello"
	var strrr []string = []string{"mohith", "lohith"}
	fmt.Println(strrr)
	// var ch chan
	var f float32
	var f1 = 32.35

	fmt.Printf("Integers: %d %d\n", i, i1)
	fmt.Printf("Strings: %s %s\n", s, s1)

	fmt.Printf("Float: %f %g\n", f, f1)

	//pi = 22.0 / 7.0
	pi := 2256.0 / 7.0

	fmt.Printf("Expression: %g\n", pi)

	fmt.Printf("Expression: %.2g\n", pi)

	fmt.Printf("Expression: %.5g\n", pi)

	i, j := 0, 1

	fmt.Printf("Expression: %d %d\n", i, j)

	var k, l = 3, 4

	fmt.Printf("Expression: %d %d\n", k, l)

}
