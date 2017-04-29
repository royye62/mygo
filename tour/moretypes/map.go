package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

// Map literals
var m2 = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967}, // you can omit type name Vertex
    "Google":    {37.42202, -122.08408},
}

func main() {
    fmt.Println(m2)

    m = make(map[string]Vertex)
    fmt.Println(m)

    // Insert or update an element
    m["Bell Labs"] = Vertex {
        40.68433, -74.39967,
    }
    m["Google"] = Vertex {
        37.42202, -122.08408,
    }
    fmt.Println(m)

    // Retrieve an element
    elem := m["Bell Labs"]
    fmt.Println(elem)

    // Delete an element
    delete(m, "Bell Labs")
    fmt.Println(m)

    // Test that a key is present with a two-value assignment
    elem, ok := m["Bell Labs"] 
    fmt.Println(elem, ok)
}


