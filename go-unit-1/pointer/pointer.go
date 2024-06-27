// package main

// import "fmt"

// func main() {

//     x := 1
//     y := 0

//     p := &x // p, of type *int, points to x

//     fmt.Println(*p) // Prints the value pointed to by p, which is the value of x

//     *p = 2 // Value stored at the address is being changed to 2 without referencing the variable x

//     fmt.Println(x) // Prints the updated value of x

//     q := &x

//     fmt.Println(p == q) // Prints true, as p and q point to the same variable x

//     q = &y

//     fmt.Println(p == q) // Prints false, as p and q now point to different variables x and y

//     fmt.Println(incr(&y)) // Calls the incr function with the address of y, prints the incremented value of y

//     fmt.Println(incr(q)) // Calls the incr function with the address of y, prints the incremented value of y

//     l := new(int) // l, of type *int, points to an unnamed int variable

//     fmt.Println(*l) // Prints the initial value pointed to by l, which is 0

//     *l = 2 // Sets the unnamed int variable to 2

//     fmt.Println(*l) // Prints the updated value pointed to by l, which is now 2

//     l1 := *l // named int variable l1 initialized to the value pointed by l

//     fmt.Println(l1) // Prints the value of l1, which is 2

// }

// func incr(p *int) int {
//     *p++ // Increments the value pointed to by p
//     return *p // Returns the incremented value
// }


package main

import (
    "fmt"

)

func main(){
    x:=1
    y:=0
    p:=&x

    fmt.Println(*p)
     *p=2;
     fmt.Println(x)
     q:=&x

     fmt.Println(p==q)
     q=&y

     fmt.Println(p==q)

    fmt.Println(incre(&y))
    fmt.Println(incre(q))

    l:=new(int)
    fmt.Println(*l)
    *l=345
    fmt.Println((*l))
    l1:=l
    fmt.Println(l1)
    fmt.Println(&l)
    fmt.Println(l)
    fmt.Println(q)
    fmt.Println(&y)
    fmt.Println(*l1)
    // fmt.Println(**l1)
}
func incre(p *int) int{
    *p++;
    return *p
}