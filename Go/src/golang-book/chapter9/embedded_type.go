package main 

import "fmt"

type Person struct {
  Name string
}

func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

// Use the type "Person and don't give it a name". This is saying an android is a person
type Android struct {
  Person
  Model string
}

// accessing the embedded Person struct through the Android Struct
func main() {
  a := new(Android)
  a.Talk()
}



