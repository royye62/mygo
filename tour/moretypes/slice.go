package main

import (
    "fmt"
)

func main() {
    // Slices are like references to arrays
    // Changing the elements of a slice modifies the corresponding elements of its underlying array.

    // slice literals
    // A slice literal is like an array literal without the length. And this creates the same array as above, then builds a slice that references it
    s1 := []struct{
        i int
        b bool
    }{
        {2, true},
        {3, false},
        {5, true},
        {7, true},
        {11, false},
        {13, true},
    }
    fmt.Println(s1)

    // A slice has both a length and a capacity
    // The length of a slice is the number of elements it contains.
    // The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice
    s := []int{2, 3, 5, 7, 11, 13}
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
    s = s[:0]
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
    s = s[:4]
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
    s = s[2:]
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

    // The zero value of a slice is nil
    // A nil slice has a length and capacity of 0 and has no underlying array
    var s2 []int
    fmt.Println(s2, len(s2), cap(s2))
    if s2 == nil {
        fmt.Println("nil!")
    }

    // Creating a slice with make
    pow := make([]int, 10) // len(pow)=10
    // pow := make([]int, 5, 10) // len(pow)=5, cap(pow)=10

    // The range form of the for loop iterates over a slice or map.
    // When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
    for i := range pow {
        pow[i] = 1 << uint(i) // == 2*i 
    }
    for _, v := range pow {
        fmt.Printf("%d\n", v)   
    }
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }

    // appending a slice
    pow = append(pow, 1111)
    printSlice(pow)

    pow = append(pow, 11, 22, 33, 44, 55, 66, 77, 88, 99, 00)
    printSlice(pow)
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
