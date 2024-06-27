// package main

//	func main() {
//		var a [3]int
//
// var str []string=[]string{""}
//
//		var b [4]int = [4]int{1, 2, 3, 4}
//		fmt.Println(a, b)
//		// var c[5]int=[5]int{1,2,3,4}
//		c := [...]int{1, 2}
//		println(c[1] + c[0])
//		k := [...]int{99: -1}
//		fmt.Println(k[99], k[98])
//		j := [...]string{99: "mohith"}
//		fmt.Println(j[99])
//	}
// type movie struct {
// 	Title  string
// 	Year   int
// 	Color  bool
// 	Actors []string
// }

//	var movies =[]movie{
//		{Title: "mo",Year:1444,Color:true,Actors: []string{}},
//		{}
//	}
// func main() {
// 	// data, err := json.Marshal(movies)
// 	var b bool = true
// 	fmt.Println(b)
// 	c := bool(false)
// 	// c = true
// 	fmt.Println(c)

// 	mapping := map[string]int{
// 		"mohith": 17,
// 		"saroja": 47,
// 	}
// 	fmt.Println(mapping)

// }

// func main(){
// 	months:=[...]{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

// }

// Q2:=months[4:7]
// endlesssummer:=Q2[:5]

// m1 := []string{"apple", "banana", "peach"}
// 	m2 := []string{"orange", "pineapple", "grapes"}

// 	m:=append(m1,m2...)

// type geometry interface{
// 	area() float64
// 	perim() float64
// }

// type rect struct {
//     width, height float64
// }

// type circle struct {
//     radius float64
// }

// func (r rect) area() float64{
// 	return r.width*r.height
// }

// func (r rect) perim() float64{
// 	return 2*(r.width + r.height)
// }

// func

// func main(){

// }

// func main() {
// 	var i interface{} = "mohith"

// 	str := i.(string)
// 	fmt.Println(str)

// 	strr, whats := i.(int)
// 	fmt.Println(strr, whats)
// }

// type ByteCounter int

// func (c *ByteCounter) Write(p []byte) (int, error) {
// 	*c += ByteCounter(len(p))
// 	return len(p), nil
// }

// func main() {
// 	r := strings.NewReader("hellocjwsnviwbnvow")

// 	b := make([]byte, 8)

// 	for {
// 		n, err := r.Read(b)

//			fmt.Printf("%s\n\n", r)
//			fmt.Printf("n=%v err=%v b=%v\n\n", n, err, b)
//			fmt.Printf("b[:n] = %q\n", b[:n])
//			if err == io.EOF {
//				break
//			}
//		}
//	}
// func comparing(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Printf("Twice %v is %v\n", v, v*2)
// 	case string:
// 		fmt.Printf("%q is %v bytes long\n", v, len(v))
// 	case float64:
// 		fmt.Printf("Twice %v is %v\n", v, v*2)
// 	default:
// 		fmt.Printf("I don't know about type %T!\n", v)
// 	}
// }
// func main() {

// }

// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	// fmt.Println(<-ch)

// }

// func summerchannel(s []int, c chan int) {
// 	sum := 0
// 	for _, val := range s {
// 		sum += val
// 	}
// 	c <- sum
// }

// func main() {
// 	var arr []int = []int{7, 2, 8, -9, 4, 0}
// 	c := make(chan int)

// 	go summerchannel(arr[3:6], c)
// 	go summerchannel(arr[0:3], c)

// 	x := <-c
// 	y := <-c
// 	fmt.Println(x, y, x+y)
// 	close(c)

// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	naturals := make(chan int)
// 	squares := make(chan int)

// 	go func() {
// 		for x := 0; x < 10; x++ {
// 			naturals <- x
// 		}
// 		close(naturals)
// 	}()

// 	go func() {
// 		for x := range naturals {
// 			squares <- x * x
// 		}
// 		close(squares)
// 	}()

// 	for x := range squares {
// 		fmt.Println(x)
// 	}
// }

// package main

// import (
// 	"time"
// 	"fmt"
// )

// func main(){
// 	go spinner(100*time.Millisecond)
// 	const n=39
// 	fibN:=fib(n)
// 	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
// 	//fmt.Printf("\r%d\n", n)
// }
// func fib(x int) int {
// 	if x < 2 {
// 		return x
// 	}

// 	return fib(x-1) + fib(x-2)
// }

// func spinner(delay time.Duration){
// 	for {
// 		for _,x:=range '\|/'{
// 			fmt.Printf("\r%c",x)
// 			time.Sleep(delay)
// 		}
// 	}
// }

// package main

// func main() {
// 	naturals := make(chan int)
// 	squares := make(chan int)
// 	go counter(naturals)
// 	go squarer(naturals, squares)

// }

// func counter(out chan<- int) {
// 	for x := 0; x < 10; x++ {
// 		out <- x
// 	}
// 	close(out)
// }

// func squarer(in <-chan int, out chan<- int) {
// 	for v := range in {
// 		out <- v * v
// 	}
// 	close(out)
// }

// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// 	"time"
// )

// func main() {
// 	var counter int64

// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			atomic.AddInt64(&counter, 1)
// 			fmt.Println("Incremented : ", atomic.LoadInt64(&counter))
// 			time.Sleep(time.Millisecond)
// 		}
// 	}()

// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			atomic.AddInt64(&counter, -1)
// 			fmt.Println("Decremented : ", atomic.LoadInt64(&counter))
// 			time.Sleep(time.Millisecond)
// 		}
// 	}()

// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Final value of counter: ", atomic.LoadInt64(&counter))
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"sync/atomic"
// )

// func main() {
// 	var ops atomic.Uint64
// 	var wg sync.WaitGroup

// 	for i := 0; i < 50; i++ {
// 		wg.Add(1)

// 		go func() {
// 			for c := 0; c < 1000; c++ {
// 				ops.Add(1)

// 			}
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(ops.Load())
// }

// package main

// import (
// 	"sync"
// 	"time"
// )

// type SafeCounter struct {
// 	mu sync.Mutex
// 	v  map[string]int
// }

// func (c *SafeCounter) Inc(key string) {
// 	c.mu.Lock()

// 	c.v[key]++
// 	defer c.mu.Unlock()

// }
// func (c *SafeCounter) Value(key string) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.v[key]++
// }

// func main() {
// 	c := SafeCounter{v: make(map[string]int)}
// 	for i := 0; i < 1000; i++ {
// 		go c.Inc("somekey")
// 	}

// 	time.Sleep(time.Second)
// 	// fmt.Println(c.Va)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var GFG int

// func worker(wg *sync.WaitGroup, m *sync.Mutex) {

// 	m.Lock()
// 	GFG = GFG + 1

// 	m.Unlock()
// 	wg.Done()
// }

// func main() {
// 	var m sync.Mutex
// 	var w sync.WaitGroup

// 	for i := 0; i < 1000; i++ {
// 		w.Add(1)
// 		go worker(&w, &m)
// 	}

// 	w.Wait()
// 	fmt.Println(GFG)
// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type mobile struct {
// 	price float64
// 	color string
// }

// func main() {
// 	s1 := []string{"A", "B", "C", "D", "E"}
// 	s2 := []string{"D", "E", "F"}
// 	result := reflect.DeepEqual(s1, s2)
// 	fmt.Println(result)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var val float64 = 15
// 	value2 := "Go programming language"
// 	observe(val)
// 	observe(value2)
// }

// func observe(i interface{}) {
// 	fmt.Println("The type passed is : %T\n", i)
// 	fmt.Printf("The value passed is: %#v \n", i)
// 	fmt.Println("-------------------------------------")
// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type First struct {
// 	A int
// 	B string
// 	C float64
// }

// type Second struct {
// 	First
// 	D bool
// }

// func main() {
// 	var temp First
// 	temp = First{10, "ABCD", 15.20}
// 	s := Second{temp, true}
// 	t := reflect.TypeOf(s)
// 	fmt.Printf("%v\n", t.FieldByIndex([]int{0}))
// 	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 0}))
// 	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 1}))
// 	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 2}))
// 	fmt.Printf("%v\n", t.FieldByIndex([]int{1}))

// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type T struct {
// 	A int
// 	B string
// 	C float64
// }

// func main() {
// 	s := T{10, "abcd", 15.20}
// 	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("A"))
// 	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("B"))
// 	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("C"))

// 	reflect.ValueOf(&s).Elem().FieldByName("B").SetString("Test")
// }

package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type T struct {
// 	A int
// 	B string
// 	C float64
// 	D bool
// }

// func main() {
// 	// destination := reflect.ValueOf([]string{"A", "B", "E"})
// 	// source := reflect.ValueOf([]string{"D", "E", "F"})
// 	// fmt.Println(destination)
// 	// counter := reflect.Copy(destination, source)
// 	// fmt.Println(counter)
// 	// fmt.Println(source)
// 	// fmt.Println(destination)

// 	t := T{10, "ABCD", 15.20, true}
// 	typeT := reflect.TypeOf()

// 	for i := 0; i < typeT.NumField(); i++ {
// 		field := typeT.Field(i)
// 		fmt.Println(field.Name, field.Type)

// 	}
// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main(){
// 	theList:=[]int{1,2,3,4}
// 	swap:=reflect.Swapper(theList)
// 	swap(0,3)

// 	for i:=i<len(theList)/2;i++{
// 		swap(i,len(theList)-1-i)

// 	}
// }


package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
)


func handler(w http.ResponseWriter,r *http,Request){
	var i =false
	var j=false
	var k=false

	searchkeyword:=r.URL.Query().Get("word")
}


func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}