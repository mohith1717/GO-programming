package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
	C float64
}

func main() {
	s := T{10, "ABCD", 15.20}
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("A"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("B"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("C"))

	reflect.ValueOf(&s).Elem().FieldByName("A").SetInt(50)
	reflect.ValueOf(&s).Elem().FieldByName("B").SetString("Test")
	reflect.ValueOf(&s).Elem().FieldByName("C").SetFloat(5.5)

	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("A"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("B"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("C"))
}
