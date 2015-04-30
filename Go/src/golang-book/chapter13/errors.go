package main 

import (
  "fmt"
  "errors"
)

func main() {
  err := errors.New("error message")
  fmt.Println(err)
}