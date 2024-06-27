package main

import "fmt"

func main() {

	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	Q2 := months[4:7]

	summer := months[6:9]

	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"] June is included in each and is the sole output of this (inefficient)

	// Extending a slice
	endlessSummer := Q2[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)

	// Appending slices

	m1 := []string{"apple", "banana", "peach"}
	m2 := []string{"orange", "pineapple", "grapes"}

	m := append(m1, m2...)
	m = append(m, "mohith")
	fmt.Println(m)

	//A slice can be used to implement a stack. Given an initial ly empty slice stack, we can push a new value onto the end of the slice with append:

	var stack []string
	v := "apple"

	stack = append(stack, v) // push v

	fmt.Println(stack)

	//The top of the stack is the last element:
	top := stack[len(stack)-1]

	fmt.Println(top)
	v2 := "orange"
	stack = append(stack, v2)
	fmt.Println(stack)
	stack = stack[:len(stack)-1] // pop

	fmt.Println(stack)
}
