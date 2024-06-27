package main

import "fmt"

func main() {

	fmt.Println(sum())

	fmt.Println(sum(1, 2, 3))

	fmt.Println(sum(1, 2, 3, 4, 5))
}

func sum(vals ...int) int {
	total := 0

	for _, val := range vals {
		total += val
	}
	return total
}

func somer(vall ...int) int {
	tot := 0
	for _, val := range vall {
		tot += val
	}
	return tot
}