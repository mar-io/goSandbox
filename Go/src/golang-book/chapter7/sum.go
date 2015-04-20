package main 

import "fmt"

// create the functions signature by defining the parameters and the their type and the functions return type
func sum(xs []int) int {
  //panic is built-in function that causes run-time error
  //panic("Not Implemented")
  total := 0
  for _, v := range xs {
    total += v
  }
  return total
}

func main() {
  // name of parameters don't have to match those called in the average function. So 'xs' can be any variable.
  xs := []int{98,93,77,82,83}
  fmt.Println(sum(xs))
}