package main

import "fmt"

func main() {

	var f32 float32 = 16777216

	var f64 float64 = 16777216 * 2

	const Avogadro = 6.02214129e23

	fmt.Printf("%f\n %f\n %f\n", f32, f64, Avogadro)

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

}
