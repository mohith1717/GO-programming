package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	var x complex128 = complex(1, 2) // 1+2i
	var m complex64 = complex(34, 3)

	var y complex128 = complex(3, 4) // 3+4i

	fmt.Println(x) // "(1+2i)"
	fmt.Println(y) // "(3+4i)"
	fmt.Println(m)
	fmt.Println(x * y)       // "(5+10i)"
	fmt.Println(real(x * y)) // "5"
	fmt.Println(imag(x * y)) // "10"

	fmt.Println(x == y)
	fmt.Println(x != y)

	fmt.Println(cmplx.Sqrt(1))

}
