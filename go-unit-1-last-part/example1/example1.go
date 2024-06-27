package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(hypot(3, 4))

	fmt.Println(first(3, 4))
}

func hypot(x, y float64) float64 { return math.Sqrt(x*x + y*y) }

func first(x int, _ int) int { return x }
