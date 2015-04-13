package main 

import "fmt"

func main() {
  x := [6]string{"a","b","c","d","e","f"}
  fmt.Println(x[2:5])
}

//should get c,d,e as the low:high index of 2:5 should give (starting at 0) c and the high index of 5 ( ends at f but doe not include f) outputs e