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

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0})) //%#v prints a value
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 2}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{1}))
}

/*
OUTPUT:
reflect.StructField{Name:"First", PkgPath:"", Type:(*reflect.rtype)(0x8e65a0), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}:

Name: The name of the field, which is "First".
PkgPath: The package path of the field. Since it's an empty string, it means it's in the same package.
Type: The reflection type of the field. In this case, it's a pointer to the reflection type of First.
Tag: The struct tag associated with the field. It's empty here.
Offset: The offset within the struct where this field begins. Since it's an embedded field, the offset is 0.
Index: The field index within the struct type hierarchy. In this case, it's [0], indicating it's the first field.
Anonymous: Indicates whether the field is anonymous (embedded). Here, it's true.*/
