package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func sayHello(wg *sync.WaitGroup, name string) {
	mutex.Lock()         // Acquire the lock before accessing and modifying the counter
	defer mutex.Unlock() // Release the lock after accessing and modifying the counter
	counter++
	fmt.Println("Hello,", name, " - Counter:", counter)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Add 2 goroutines to the wait group

	go sayHello(&wg, "Alice")
	go sayHello(&wg, "Bob")

	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("All greetings sent.")
}

/*•	var counter int: This line declares a global variable named counter of type int. It's used to keep track of the number of times the sayHello function is called. Since it's declared outside any function, it's accessible from all functions in the program.
•	var mutex sync.Mutex: This line declares a global variable named mutex of type sync.Mutex. A mutex is a synchronization primitive used to control access to shared resources. In this program, it will be used to ensure that only one goroutine can modify the counter variable at a time.
•	func sayHello(wg *sync.WaitGroup, name string) {: This line defines a function named sayHello. It takes two arguments:
o	wg *sync.WaitGroup: This is a pointer to a sync.WaitGroup object. This wait group will be used to synchronize the completion of multiple goroutines.
o	name string: This is a string argument that will be used to personalize the greeting message.
•	mutex.Lock(): This line acquires the lock on the mutex object. This ensures that only one goroutine can execute the code between Lock() and Unlock() at a time. Any other goroutine trying to acquire the lock will be blocked until the current goroutine releases it.
•	defer mutex.Unlock(): This line uses the defer statement to ensure that the mutex.Unlock() call is executed even if there's an error or panic within the function. This guarantees that the lock is always released, preventing deadlocks.
•	counter++: This line increments the counter variable. However, since it's a shared variable accessed by multiple goroutines, the mutex ensures that this increment happens atomically (as a single, indivisible operation), preventing race conditions.
•	fmt.Println(...): This line uses the fmt.Println function from the imported package to print a greeting message with the provided name and the current value of the counter.
•	wg.Done(): This line signals to the wg (wait group) that the current goroutine (the execution of the sayHello function) has finished its work.
•	func main() {: This line defines the main function, which is the entry point for program execution.
•	var wg sync.WaitGroup: This line declares a local variable named wg of type sync.WaitGroup. It's used to keep track of the number of goroutines that need to finish before the main function continues execution.
•	wg.Add(2): This line increments the counter of the wait group by 2. This tells the wait group that it needs to wait for 2 goroutines to finish their work before continuing.
•	go sayHello(&wg, "Alice"): This line launches a new goroutine (lightweight thread) that executes the sayHello function. It passes a reference to the wg (wait group) and the string "Alice" as arguments to the function.
*/
