package main

import "fmt"

func main() {

	var b1 bool = true

	var b2 bool = false

	fmt.Println(b1)
	fmt.Println(b1)
	fmt.Println(b1 == b2)
	fmt.Println(b1 != b2)

	var i int

	//  i = int(b1)
	if b1 {
		i = 11
	}

	fmt.Println(i)

	//i = int(b2)
	if !b2 {
		i = 10
	}

	fmt.Println(i)

}
