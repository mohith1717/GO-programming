package main

import "fmt"

// "math/cmplx"

func main() {
	// 	var arr [3]int
	// 	fmt.Println(arr[0])
	// 	fmt.Println(arr[len(arr)-1])

	// 	var aq [3]int = [...]int{2,3,4}
	// 	fmt.Println(aq[len(aq)-1])

	// 	k:=[...]int{28:99}
	// 	fmt.Println(k[len(k)-1])
	// 	var x complex128 = complex(1, 2)
	// 	var y complex128 = complex(3, 4)
	// 	fmt.Println(x * y)
	// 	fmt.Println(real(x * y))
	// var pi = 22.0/7.0
	// var i8 uint8

	// i8 = uint8(pi)
	// fmt.Println(i8)

	// var str string
	// str = "hello, world"
	// fmt.Println(str)
	// fmt.Println(str[0],str[7])

	// var x string
	// x = "1234"
	// y,err := strconv.Atoi(x) //converts string into integer̦

	// fmt.Println(y+2,err)

	// z:=strconv.Itoa(y)
	// fmt.Println(z+"1")

	// var some string
	// some = "123.456"
	// something,err :=strconv.ParseFloat(some,32)
	// fmt.Println(something,err)

	// someone,err :=strconv.ParseFloat(some,64)
	// fmt.Println(someone,err)

	// fmt.Println(divide(10, 2))

	// fmt.Println(divide(10, 0))

	// f:=sqaures()
	// fmt.Println(f())

	// fmt.Println(f())

	// fmt.Println(f())
	// fmt.Println(f())

	// calculating the mode
	// q := []int{1, 1, 2, 2, 3, 3, 3, 3, 4, 4}
	// res := mode(q)
	// fmt.Println(res)

	var i int
	i = 23
	var pointer *int
	pointer = &i
	fmt.Println(&i)
	fmt.Printf("Value of i is %d\n", i)
	fmt.Println(pointer)
	var _2pointer *int
	_2pointer = pointer
	fmt.Println(_2pointer)
	fmt.Println(*_2pointer)
}

// func mode(a []int) int {
// 	var actualmode int
// 	actualmode = a[0]
// 	var actualfreq int
// 	actualfreq = 0

// 	for i := 0; i < len(a); {
// 		runlen := 1
// 		runval := a[i]

// 		for j := i + 1; j < len(a) && a[j] == runval; j++ {
// 			runlen++
// 		}

// 		if runlen > actualfreq {
// 			actualfreq = runlen
// 			actualmode = runval
// 		}
// 		i += runlen
// 	}

// 	return actualmode
// }
// func divide(a, b int) int {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("In recovery ", r)
// 		}
// 	}()
// 	if b == 0 {
// 		panic("Divide by zero")

// 	}
// 	if b != 0 {
// 		return a / b
// 	}
// 	return 0
// }

// func fact(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * fact(n-1)
// }
