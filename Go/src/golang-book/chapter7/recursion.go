package main 

import "fmt"

func factorial(x uint) uint {
  if x == 0 {
    return 1
  }
  // factorial calling itself making it RECURSIVE.
  return x * factorial(x-1)
}

// So for example factorial(2)
// Is x == 0 no x is 2
// return 2 * factorial(2-1)
// // Is x == 0 no x is 1
// // return 1 * factorial(1-1)
// // // Is x == 0 yes. return 1
// // // return 1 * 1 = 1
// return 2 * 1