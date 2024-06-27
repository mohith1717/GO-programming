package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [4]int{1, 2, 3, 4}
	ptr := &arr[0]

	// Move the pointer to the next element in the array
	nextPtr := unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0]))
	fmt.Println("Value of the next element:", *(*int)(nextPtr))
}
