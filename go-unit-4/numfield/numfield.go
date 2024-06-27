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
     fmt.Println(typeT)
	fmt.Println(typeT.NumField())
}

