package main 

import ("fmt"; "math")

// similar methods use similar interfaces
// define a method set or a list of methods that a type must have in order to implement the interface

type Shape interface {
  area() float64
}
// since recatangle and circle both have area methods that return float64
// both types implement the Shape interface

  func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
      area += s.area()
    }
    return area
  }

type Rectangle struct {
// collapsing x,y,r in the same field since they all have float64 type
  x1,y1,x2,y2 float64
}

  func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
  }

  func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
  }

type Circle struct {
  x,y,r float64
}

  func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
  }

func main() {
  r := Rectangle{0, 0, 10, 10}
  c := Circle{0, 0, 5}
  fmt.Println(totalArea(&c, &r))
}





