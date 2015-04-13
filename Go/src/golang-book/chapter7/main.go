package main 

import "fmt"

// create the functions signature by defining the parameters and the their type and the functions return type
func average(xs []float64) float64 {
  //panic is built-in function that causes run-time error
  //panic("Not Implemented")
  total := 0.0
  for _, v := range xs {
    total += v
  }
  return total / float64(len(xs))
}

func main() {
  // name of parameters don't have to match those called in the average function. So 'xs' can be any variable.
  xs := []float64{98,93,77,82,83}
  fmt.Println(average(xs))
}