package main

import (
	"fmt"
	"io"
	"strings"
)

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// 	j int
// }

// func (t T) M() {
// 	fmt.Println(t.S)
// 	fmt.Println(t.j)
// }
type geometry interface{
	area()
}

func main() {
// 	var i I = T{"Hello", 56}
// 	i.M()
	r:=strings.NewReader("Hello students")

	b:=make([]byte,8)

	for{
		n,err := r.Read(b)

			fmt.Println("%s",r)
		fmt.Println("n=")

	}
}
