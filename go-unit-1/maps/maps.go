package main

import "fmt" 

func main() {

     ages := map[string]int{
                             "alice": 31,
                             "charlie": 34,
                           }

     fmt.Println(ages["alice"])

     delete(ages, "alice") // remove element ages["alice"]

     fmt.Println(ages["alice"])

     ages["charlie"]++

     fmt.Println(ages["charlie"])

     ages["charlie"]--

     fmt.Println(ages["charlie"])

}
