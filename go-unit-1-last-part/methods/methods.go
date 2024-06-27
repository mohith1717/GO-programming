package main

import "fmt"

type rect struct {
	width, height int
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", area(r))
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", area(*rp))
	fmt.Println("perim:", rp.perim())
}

// func (r rect) area() int {
// 	return r.width * r.height
// }

// func area(r rect) int {
// 	return r.width * r.height

// }

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// package main

// import "fmt"

// func main() {

// 	var x int = 5

// 	f2(x)
// 	a, b, c := f1(&x) // pass by reference

// 	// fmt.Println(some)
// 	fmt.Println(a, b, c)
// }

// func f2(x int) {
// 	fmt.Println(x)
// }
// func f1(x *int) (r1, r2, r3 int) { //var int r3
// 	//r3 = r1 * r2
// 	r3 = *x * 9
// 	*x = *x - 1
// 	r2 = *x - 1
// 	r1 = *x + 1

// 	return // bare return
// }
