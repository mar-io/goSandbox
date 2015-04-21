package main 

import "fmt"

// * character represents pointer follwed by stored value
func square(x *float64) {
  // pointers reference a location in memory where a valued is stored
  // store the int 0 in the memory location xPtr refers to
  *x = *x * *x
}

func main() {
  x := 1.5
  // & operator to find the address of the variable. eseet returns an *int because x in an int.
  square(&x) // x should be 2.25
  fmt.Println(x)
}