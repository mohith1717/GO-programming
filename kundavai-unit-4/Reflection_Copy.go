package main

import (
	"fmt"
	"reflect"
)

func main() {
	destination := reflect.ValueOf([]string{"A", "B", "C"})
	source := reflect.ValueOf([]string{"D", "E", "F"})

	// Copy() function is used and it returns the number of elements copied
	counter := reflect.Copy(destination, source)
	fmt.Println(counter)

	fmt.Println(source)
	fmt.Println(destination)
}