package main

import "fmt"

// A type assertion provides access to an interface value's underlying concrete value. 
// t := i.(T)
// t, ok := i.(T)

func main() {
    var i interface{} = "hello"

    // asserts s implements string
    s := i.(string)
    fmt.Println(s)

    s, ok := i.(string)
    fmt.Println(s, ok)

    f, ok := i.(float64)
    fmt.Println(f, ok)

    //f, ok := i.(float64)  // no new variables on left side of :=   f, ok must have one is not declare
    f1 := i.(float64) // panic
    fmt.Println(f1, ok)
}
