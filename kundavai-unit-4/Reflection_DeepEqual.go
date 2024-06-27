package main

import (
	"fmt"
	"reflect"
)

type mobile struct {
	price float64
	color string
}

func main() {
	// DeepEqual is used to check two slices are equal or not
	s1 := []string{"A", "B", "C", "D", "E"}
	s2 := []string{"D", "E", "F"}
	result := reflect.DeepEqual(s1, s2)
	fmt.Println(result)

	// DeepEqual is used to check two arrays are equal or not
	n1 := [5]int{1, 2, 3, 4, 5}
	n2 := [5]int{1, 2, 3, 4, 5}
	result = reflect.DeepEqual(n1, n2)
	fmt.Println(result)

	// DeepEqual is used to check two structures are equal or not
	m1 := mobile{500.50, "red"}
	m2 := mobile{400.50, "black"}
	result = reflect.DeepEqual(m1, m2)
	fmt.Println(result)
}