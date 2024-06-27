package main

import "fmt"

func main() {

	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	i8 = -1
	i16 = -1
	i32 = -1
	i64 = -1

	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64

	// var temp int8 = -1
	// u8 = -1 * -1
	u8 = -1 * -1
	var uu uint32
	uu = 1
	var uu2 uint32
	uu2 = uu - (4)

	fmt.Println(uu2)
	fmt.Println(uu)
	u16 = -1 * -1
	u32 = -1 * -1
	u64 = -1 * -1

	//i16 = i8

	i16 = int16(i8)

	fmt.Printf("Signed Integers: %d %d %d %d\n", i8, i16, i32, i64)

	fmt.Printf("Unsigned Integers: %d %d %d %d\n", u8, u16, u32, u64)

	var pi = 22.0 / 7.0
	var some float64
	some = float64(pi)

	fmt.Printf("%f %f\n", pi, some)

	var k int

	k = int(pi)

	fmt.Printf("%f %d\n", pi, k)

}
