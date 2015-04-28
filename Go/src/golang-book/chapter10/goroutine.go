package main

import (
    "fmt"
    "time"
    "math/rand"
)

func f(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
    // will create amt with random amount between 0 and 250
    amt := time.Duration(rand.Intn(250))
    // will wait between 0 and 250 ms between each goroutine
    time.Sleep(time.Millisecond * amt)
  }
}

// iterate multiple go routines, the first goroutine is implicit and second is when we call go f(i)
func main() {
  for i := 0; i < 10; i++ {
    go f(i)
  }
  var input string
  // scanln function included so that programs lets the goroutine run before it exits
  fmt.Scanln(&input)  
}