package main 

import "fmt"

// * character represents pointer follwed by stored value
func one(xPtr *int) {
  // pointers reference a location in memory where a valued is stored
  // store the int 0 in the memory location xPtr refers to
  *xPtr = 1
}

func main() {
  // new is a builtin function that takes type as an argument and allocates enough memory to fit a value of that type
  xPtr := new(int)
  one(xPtr)
  // & operator to find the address of the variable. eseet returns an *int because x in an int.
  fmt.Println(*xPtr) // x is 0
}

// new and & are not different since go is garbage collected and memory is clean up when nothing refers to it