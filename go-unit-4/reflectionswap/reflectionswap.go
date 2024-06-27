package main

import (
	"fmt"
	"reflect"
)

func main() {
	theList := []int{1, 2, 3, 4, 5}
	swap := reflect.Swapper(theList)
	fmt.Printf("Original Slice :%v\n", theList)

	// Swapper() function is used to swaps the elements of slice
	swap(0, 3)
	fmt.Printf("After Swap :%v\n", theList)

	// Reversing a slice using Swapper() function
	for i := 0; i < len(theList)/2; i++ {
		swap(i, len(theList)-1-i)
	}
	fmt.Printf("After Reverse :%v\n", theList)
}
