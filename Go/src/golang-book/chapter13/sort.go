package main 

import (
  "fmt"
  "sort"
)

type Person struct {
  Name string
  Age int
}
// to define our own sort we create a new type and make it equivalent to a slice we want to sort
type ByName []Person
// sort interface requires 3 methods Len, Less, Swap. Cast our list into a type.
func (this ByName) Len() int {
  return len(this)
}

func (this ByName) Less(i, j int) bool {
  return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
  this[i], this[j] = this[j], this[i]
}

func main() {
  kids := []Person{
    {"Jill",9},
    {"Jack",10},
  }
  sort.Sort(ByName(kids))
  fmt.Println(kids)
}