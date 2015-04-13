package main 

import "fmt"

func main() {
    fmt.Println("Enter a how many feet: ")
    var feet float64
    fmt.Scanf("%f", &feet)
    
    meters := feet * 0.3048

    fmt.Println(meters)
}