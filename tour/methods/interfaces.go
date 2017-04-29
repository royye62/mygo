package main

import (
    "fmt"
    "math"
)

// interface values can be thought of as a tuple of a value and a concrete type: (value, type)
// Calling a method on an interface value executes the method of the same name on its underlying type.

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    // it is common to write methods that gracefully handle being called with a nil receiver
    if t == nil {
        fmt.Println("<nil>") 
        return
    }
    fmt.Println(t.S)
}

type F float64

func (f F) M() {
    fmt.Println(f)
}

func main() {
    var i I

    i = &T{"Hello"}
    describe(i)
    i.M()

    i = F(math.Pi)
    describe(i)
    i.M()

    /// Interface values with nil underlying values
    var t *T
    i = t // an interface value that holds a nil concrete value is itself non-nil
    describe(i)
    i.M()

    /// Nil interface values
    var i2 I
    describe(i2) // A nil interface value holds neither value nor concrete type.
    // i2.M() // Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

    /// The empty interface
    // The interface type that specifies zero methods is known as the empty interface: interface{}
    var i3 interface{}
    describe0(i3)
    // An empty interface may hold values of any type. (Every type implements at least zero methods.)
    i3 = 42
    describe0(i3)
    i3 = "hello"
    describe0(i3)
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}

func describe0(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

