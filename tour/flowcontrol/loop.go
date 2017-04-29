package main

import "fmt"

// go has only one looping construct, the for loop

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i 
    }
    fmt.Println(sum)


    sum1 := 1
    for sum1 < 1000 {
        sum1 += sum1
    }
    fmt.Println(sum1)
}

