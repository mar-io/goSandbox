package main 

import "fmt"

// * character represents pointer follwed by stored value
func zero(xPtr *int) {
  // pointers reference a location in memory where a valued is stored
  // store the int 0 in the memory location xPtr refers to
  *xPtr = 0
}

func main() {
  x := 5
  // & operator to find the address of the variable. eseet returns an *int because x in an int.
  zero(&x)
  fmt.Println(x) // x is 0
}