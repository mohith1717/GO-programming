// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"unsafe"
)

func main() {

	type IntString struct {
		I int
		S string
	}

	type StringInt struct {
		S string
		I int
	}

	type Flag struct {
		num1 int16
		num2 int32
	}

	is := IntString{}
	si := StringInt{}
	f := Flag{}

	// unsafe.Sizeof
	fmt.Println("size of is ", unsafe.Sizeof(is), "\nsize of si ", unsafe.Sizeof(si), "\nsize of f ", unsafe.Sizeof(f))

	// unsafe.Alignof
	fmt.Println("Alignment of is: ", unsafe.Alignof(is), "\nAlignment of si ", unsafe.Alignof(si), "\nAlignment of f ", unsafe.Alignof(f))
}
