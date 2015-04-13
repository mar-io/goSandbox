package main 

import "fmt"

func main() {
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
}
  min := x[0]

  for _, value := range x {
    if value < min {
      min = value
    }
  }
  fmt.Println(min)
}

//should get c,d,e as the low:high index of 2:5 should give (starting at 0) c and the high index of 5 ( ends at f but doe not include f) outputs e