package main

import "fmt"

// Constants can be character, string, boolean, or numeric values
// Constants cannot be declared using the := syntax

const Pi = 3.14

func main() {
    const World = "世界"
    const Truth = true
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")
    fmt.Println("Go rules?", Truth)

    //Truth = false // compile error. cannot assign to Truth
}
