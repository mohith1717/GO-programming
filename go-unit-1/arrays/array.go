package main

import "fmt" 

func main() {

     var a [3]int // array of 3 integers

     fmt.Println(a[0]) // print the first element

     fmt.Println(a[1]) // print the second element
     
     fmt.Println(a[len(a) - 1]) // print the last element

     var q [3]int = [3]int{1, 2, 3}

     fmt.Println(q[0]) // print the first element

     fmt.Println(q[1]) // print the second element
     
     fmt.Println(q[len(q) - 1]) // print the last element

     var r [3]int = [3]int{1, 2}

     fmt.Println(r[2]) // "0"

     k := [...]int{99: -1} 
     fmt.Println(k[0]) 
     fmt.Println(k[len(k) - 1]) // -1

     a1 := [2]int{1, 2}
     b1 := [...]int{1, 2}
     c1 := [2]int{1, 3}
     fmt.Println(a1 == b1, a1 == c1, b1 == c1) // "true false false"

}