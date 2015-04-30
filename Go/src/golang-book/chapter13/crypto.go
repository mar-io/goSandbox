package main 

import (
  "fmt"
  "hash/crc32"
)

type Person struct {
  Name string
  Age int
}

func main() {
  h := crc32.NewIEEE()
  // we can use Writer inerface in crc32
  h.Write([]byte("test"))
  // Call Sum32 to return uint32
  v := h.Sum32()
  // based on what is return if were were comparing multiple files we could deter the files are the same
  // if values are different they are definitely not the same
  fmt.Println(v)
}