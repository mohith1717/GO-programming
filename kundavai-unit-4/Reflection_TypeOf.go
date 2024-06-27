package main

import (
	"fmt"
	"reflect"
)

func main() {
	v1 := []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(v1))

	v2 := "Hello World"
	fmt.Println(reflect.TypeOf(v2))

	v3 := 1000
	fmt.Println(reflect.TypeOf(v3))

	v4 := map[string]int{"mobile": 10, "laptop": 5}
	fmt.Println(reflect.TypeOf(v4))

	v5 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(v5))

	v6 := true
	fmt.Println(reflect.TypeOf(v6))
}