package main

import (
	"fmt"
	"strconv"
)

func main() {
	var value float64 = 15
	value2 := "Go Programming Language"

	fmt.Println(Sprint(value))
	fmt.Println(Sprint(value2))
	fmt.Println(Sprint(10))
	fmt.Println(Sprint(true))
}

func Sprint(x interface{}) string {
	switch x := x.(type) {

	case string:
		return x
	case int:
		return strconv.Itoa(x)
	// ...similar cases for int16, uint32, and so on...
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map, pointer, slice, struct
		return "???"
	}
}
