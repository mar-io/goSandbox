package main 

import "fmt"

// * character represents pointer follwed by stored value
func swap(x,y *int) {
  // pointers reference a location in memory where a valued is stored
  // store the int 0 in the memory location xPtr refers to
  *x,*y = *y,*x 
}

func main() {
  x := 1
  y := 2
  // & operator to find the address of the variable. eseet returns an *int because x in an int.
  swap(&x,&y)

  fmt.Println("x =", x, "y =", y) // x is 0
}