package main

import (
	"fmt"
	"reflect"
)

func main() {
	v1 := []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.ValueOf(v1))

	v2 := "Hello World"
	fmt.Println(reflect.ValueOf(v2))

	v3 := 1000
	fmt.Println(reflect.ValueOf(v3))
	fmt.Println(reflect.ValueOf(&v3))

	v4 := map[string]int{"mobile": 10, "laptop": 5}
	fmt.Println(reflect.ValueOf(v4))

	v5 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.ValueOf(v5))

	v6 := true
	fmt.Println(reflect.ValueOf(v6))
}
