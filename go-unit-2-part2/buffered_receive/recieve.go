
package main

import "fmt"

func main() {

      ss := make(chan string, 2)

      fmt.Println(<-ss)
      fmt.Println(<-ss)

      close(ss)
}