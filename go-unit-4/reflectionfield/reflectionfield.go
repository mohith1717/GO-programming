package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
	C float64
	D bool
}

func main() {
	t := T{10, "ABCD", 15.20, true}
	typeT := reflect.TypeOf(t)

	for i := 0; i < typeT.NumField(); i++ {
		field := typeT.Field(i)
		fmt.Println(field.Name, field.Type)
	}
}
