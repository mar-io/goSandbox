package main 
// net/rpc remote procedure call exposes methods so they can be invoked over network
import (
  "fmt"
  "flag"
  "math/rand"
)

func main() {
  // Define flags name,default,passed by flag
  maxp := flag.Int("max", 6, "the max value")
  // Parse
  flag.Parse()
  // Generate a number between 0 and max
  fmt.Println(rand.Intn(*maxp))
}