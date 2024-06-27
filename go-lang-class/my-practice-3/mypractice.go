// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// 	"time"
// )

// func main() {
// 	var counter int32

// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			atomic.AddInt32(&counter, 1)
// 			fmt.Printf("Incremented: &d\n", atomic.LoadInt32(&counter))
// 			time.Sleep(time.Millisecond)
// 		}
// 	}()
// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			atomic.AddInt32(&counter, -1)
// 			fmt.Printf("Decremented: %d\n", atomic.LoadInt32(&counter))
// 			time.Sleep(time.Millisecond)
// 		}
// 	}()
// 	time.Sleep(2 * time.Second)

// 	fmt.Printf("Final Value: %d\n", atomic.LoadInt32(&counter))
// }

// Program with race condition fixed by mutex
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var GFG = 0

// func worker(wg *sync.WaitGroup, m *sync.Mutex) {
// 	m.Lock()
// 	GFG = GFG + 1
// 	m.Unlock()

// 	wg.Done()
// }

// func main() {
// 	var w sync.WaitGroup
// 	var m sync.Mutex

// 	for i := 0; i < 1000; i++ {
// 		w.Add(1)
// 		go worker(&w, &m)
// 	}
// 	w.Wait()
// 	fmt.Println("The final value of GFG is ", GFG)
// }

// package main
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type SafeCounter struct{
// 	mu sync.Mutex
// 	v map[string]int
// }

// func(c *SafeCounter)Inc(key string){
// 	c.mu.Lock()
// 	c.v[key]++
// 	c.mu.Unlock()
// }

// func(c * SafeCounter)Value(key string)int{
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.v[key]
// }

// func main(){
// 	c:=SafeCounter{v : make(map[string]int)}
// 	for i := 0; i < 1000; i++ {
// 		go c.Inc("somekey")
// 	}

// 	time.Sleep(time.Second)
// 	fmt.Println(c.Value("somekey"))
// }


package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

func main(){
	pool:=semaphore.NewWeighted(2)
	go swim("A",pool);	
	go swim("B",pool);	
	go swim("C",pool);	
	go swim("D",pool);	

}
func swim(name string,pool *semaphore.Weighted){
	log.Printf("%v I want to swim",name);
	ctx:=context.Background()

	if err:=pool.Acquire(ctx,1);
	err != nil {
		log.Printf("%v: Ops, something went wrong! I cannot acquire a lane\n", name)
		return
	}
	
}
