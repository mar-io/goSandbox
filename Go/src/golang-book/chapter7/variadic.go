package main 

import "fmt"

// The '...' means that the function can take 0 or more of the parameter
func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

func main() {
  fmt.Println(add(1,2,3))
}