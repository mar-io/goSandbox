package main 

import "fmt"

func main() {
  for i := 0; i < 100; i++ {
    if ( i%15 == 0 ) {
      fmt.Println(i, "is fizzbuzz")
    } else if ( i%3 == 0 ) {
      fmt.Println(i, "is fizz")
    } else if ( i%5 == 0 ) {
      fmt.Println(i, "is buzz")
    } else {
      fmt.Println(i, "is not divisible")
    }
  }
}