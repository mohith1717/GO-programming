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
/*
arr := [4]int{1, 2, 3, 4}: This line initializes an array named arr with four integers: 1, 2, 3, and 4.

ptr := &arr[0]: Here, we take the address of the first element of the array arr using the & operator and 
assign it to the pointer variable ptr. This means ptr now points to the memory address where the first element of the array is stored.

nextPtr := unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0])):
    a. unsafe.Pointer(ptr): This converts the pointer ptr to an unsafe pointer.
    The unsafe.Pointer type in Go allows conversion between different pointer types without any restrictions, 
    but it's considered unsafe because it bypasses the type safety of Go.
    
    b. uintptr(unsafe.Pointer(ptr)): This converts the unsafe pointer to a uintptr, which is an integer type 
    that is large enough to hold the bit pattern of any pointer. 
    This conversion allows us to perform arithmetic operations on pointers, which are not allowed directly in Go.
    
    c. unsafe.Sizeof(arr[0]): This returns the size of the element type of the array arr, which is the size of an int.
    unsafe.Sizeof returns the size in bytes, so if int is 4 bytes on your system, this will return 4.
    
    d. (uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0])): 
    This adds the size of one element of the array to the memory address pointed by ptr. 
    This effectively moves the pointer to the next element in the array.
    
    e. unsafe.Pointer(...): Finally, we convert the resulting uintptr back to an unsafe pointer. 
    Now, nextPtr points to the memory address of the next element in the array.

*(*int)(nextPtr): Here, we dereference nextPtr twice. 
The outer * dereferences the pointer, and the inner * is used to interpret the data at that memory address as an int.
So, *(*int)(nextPtr) gives us the value of the next element in the array after the one ptr originally pointed to.

fmt.Println("Value of the next element:", *(*int)(nextPtr)): Finally, we print out the value of the next element in the array.
