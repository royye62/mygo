package main

import "fmt"

var c, python, java bool
var i = 1 // If an initializer is present, the type can be omitted; the variable will take the type of the initializer
//j := 2 // error. Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

func main() {
    var i int // A var statement can be at package or function level
    fmt.Println(c, python, java, i)

    // Short variable declarations
    // Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type
    k := 3
    fmt.Println(k)
}
