package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines
func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // receive from c
    fmt.Println(x, y, x+y)

    messages := make(chan string)
    go func() { messages <- "ping" }()
    // By default sends and receives block until both the sender and receiver are ready. 
    // This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
    msg := <-messages
    fmt.Println(msg)
}
