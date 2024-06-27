package main

import "fmt"

func main() {

	var x int = 5

	r1, r2, c := f1(&x) // pass by reference

	s, t, r := someother(&x)
	fmt.Println(s, t, r)
	fmt.Println(r1, r2, c)
}

func someother(y *int) (l, m, n int) {
	m = *y - 1
	n = *y + 45
	l = *y + 1000
	return
}

func f1(x *int) (r1, r2, r3 int) {
	r1 = *x + 1
	r2 = *x - 1
	r3 = *x + 89
	return
}
