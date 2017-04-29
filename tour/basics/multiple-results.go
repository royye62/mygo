package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func test(x int) (int, int, int) { // A function can return any number of results.
    return x+1, x+2, x+3
} 

func split(sum int) (x, y int) { // named return values
    x = sum*4/9
    y = sum - x
    return // a "naked" return, should be used only in short functions, harm readability in longer functions.
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b) 

    x, y, z := test(1)
    fmt.Println(x, y, z) 

    fmt.Println(split(17))
}
