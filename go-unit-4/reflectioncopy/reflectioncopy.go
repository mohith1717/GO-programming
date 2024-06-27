package main

import (
	"fmt"
	"reflect"
)

func main() {
	destination := reflect.ValueOf([]string{"A", "B", "C"})
	source := reflect.ValueOf([]string{"D", "E", "F"})
	fmt.Println(destination)
	counter := reflect.Copy(destination, source)
	fmt.Println(counter)

	fmt.Println(source)
	fmt.Println(destination)
}
