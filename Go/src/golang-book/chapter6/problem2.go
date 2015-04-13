package main 

import "fmt"

func main() {
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
}
// sets min to first element 0 the x array of integers
  min := x[0]
// _ tells compiler we dont need iterator variable
  // value is the same as x[i]
  // range followed by the variable x we are looping over
  
  for _, value := range x {
   // if the value is less then current min then set min equal to that value
    if value < min {
      min = value
    }
  }
  fmt.Println(min)
}

