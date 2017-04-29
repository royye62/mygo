package main

import "fmt"

// One of the most ubiquitous interfaces is Stringer defined by the fmt package

type Person struct {
    Name string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func main() {
    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}
    fmt.Println(a,z)
}
