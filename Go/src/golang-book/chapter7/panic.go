package main 

import "fmt"

func main() {
  // Built in recover function is defered. This will stop the panic and return value that was passed to call panic
  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  panic("PANIC")
}