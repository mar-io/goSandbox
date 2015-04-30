package main 

import (
  "fmt"
  "container/list"
)

func main() {
  // zero value for List is an empty list. a *List can be create using list.New.
  var x list.List
  // values are appended to the list using PushBack
  x.PushBack(1)
  x.PushBack(2)
  x.PushBack(3)
  // loop over each item in the list by getting the first element and following all the links until we reach nil.
  for e := x.Front(); e != nil; e=e.Next() {
    fmt.Println(e.Value.(int))
  }
}