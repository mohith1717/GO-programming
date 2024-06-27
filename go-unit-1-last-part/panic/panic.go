package main

import "fmt"

func main() {
	// Call divide with a non-zero divisor
	fmt.Println(divide(10, 2)) // Output: 5

	// Call divide with a divisor of 0
	fmt.Println(divide(10, 0))
	// The program will panic with "Divide by zero" and exit.
	// The following line will not execute.
	fmt.Println("This line will not execute")
}

func divide(a, b int) int {
	// Check if the divisor is 0
	defer func() {
		// Recover from any panic that occurred during the execution of the deferred function.
		if r := recover(); r != nil {
			fmt.Println("In recover \n", r)
		}
	}()
	if b == 0 {
		// If b is 0, panic with a message
		panic("Divide by zero")

		// The following code is unreachable due to the panic statement above.
		// It includes a recover block to handle panics, but it will not be executed.
		if r := recover(); r != nil {
			fmt.Println("In recover ", r)
		}
	}

	// If the divisor is not 0, perform the division and return the result
	return a / b
}
