package main

import (
	"fmt"
	"reflect"
)

type First struct {
	A int
	B string
	C float64
}

type Second struct {
	First
	D bool
}

func main() {
	s := Second{First: First{10, "ABCD", 15.20}, D: true}
	t := reflect.TypeOf(s)

	fmt.Printf("%v\n", t.FieldByIndex([]int{0}))
	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 1}))
	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 2}))
	fmt.Printf("%v\n", t.FieldByIndex([]int{1}))

}
