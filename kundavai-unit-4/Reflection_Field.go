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
	//reflect.TypeOf() function is used to get
	//the reflect.Type object representing the type of t

	for i := 0; i < typeT.NumField(); i++ {
		field := typeT.Field(i)
		fmt.Println(field.Name, field.Type)
	}
}

/*NumField() is a method of the reflect.Type object.
It returns the number of fields in the struct type represented by typeT.

Field(i) is also a method of the reflect.Type object.
It returns a reflect.StructField object that represents the i-th field of the struct type. This StructField object contains information about the field, such as its name, type, tag, etc. In the provided code, it's used within the loop to retrieve each field's information.
*/
