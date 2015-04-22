package main 

import ("fmt"; "math")
// introducing a new type with name of type being "Circle" and "struct" to indicate we are defing a struct type
type Rectangle struct {
// collapsing x,y,r in the same field since they all have float64 type
  x1,y1,x2,y2 float64
}

// Access the fields using the '.' operator
//fmt.Println(c.x, c.y, c.r)
//c.x = 10
//c.y = 5

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

// This functions uses a "receiver" called "area". Allowing the use of the . operator in the main function
func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func main() {
// initialize circle with new function
// this allocates memory for all fields, sets each of them to zero, and returns a pointer(*Circle)
// c := new(Circle)
  // leaving out x:0, y:0, r:5 since I know the order the Circle struct was defined
r := Rectangle{0, 0, 10, 10}
fmt.Println(r.area())
}




