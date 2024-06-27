package main

import "fmt"

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.

// func (t T) M() {
// 	fmt.Println(t.S)
// }

func main() {
	// var i I = T{"hello"}
	// i.M()

	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(int8)
	fmt.Println(f, ok)

}
