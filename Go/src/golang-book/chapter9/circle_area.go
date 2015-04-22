package main 

import ("fmt"; "math")
// introducing a new type with name of type being "Circle" and "struct" to indicate we are defing a struct type
type Circle struct {
// collapsing x,y,r in the same field since they all have float64 type
  x,y,r float64
}

// Access the fields using the '.' operator
//fmt.Println(c.x, c.y, c.r)
//c.x = 10
//c.y = 5

func main() {
// initialize circle with new function
// this allocates memory for all fields, sets each of them to zero, and returns a pointer(*Circle)
// c := new(Circle)
  // leaving out x:0, y:0, r:5 since I know the order the Circle struct was defined
c := Circle{0, 0, 5}
fmt.Println(circleArea(&c))
}

func circleArea(c *Circle) float64 {
  return math.Pi * c.r*c.r
}
