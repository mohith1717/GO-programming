package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := "hello, world"

	fmt.Println(len(s))
	fmt.Printf("value of %c and %c\n", s[0], s[7])
	// to print ascii value
	fmt.Println(s[0], s[6])

	fmt.Println(s[0:6])

	fmt.Println(s[:5]) // "hello"
	fmt.Println(s[7:]) // "world"
	fmt.Println(s[:])  // "hello, world"

	fmt.Println(s[:5] + " students") // "hello"

	x := "123"

	y, err := strconv.Atoi(x)
	fmt.Println(y, err)

	z := strconv.Itoa(y)
	fmt.Println(z)

	x = "123.45"
	z1, err := strconv.ParseFloat(x, 32)
	fmt.Println(z1)

	z1, err = strconv.ParseFloat(x, 64)
	fmt.Println(z1)

}
