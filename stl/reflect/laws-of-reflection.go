package main

import (
    "fmt"
    "reflect"
    "io"
    "os"
)

type T struct {
A int
B string
}

func StructReflect() {
    t := T{23, "skidoo"}
    s := reflect.ValueOf(&t).Elem()
    typeOfT := s.Type()
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Println("typeOfT.Field(i):", typeOfT.Field(i))
        fmt.Printf("%d: %s %s = %v\n", i,
            typeOfT.Field(i).Name, f.Type(), f.Interface())
    }
}

func main() {
    type MyFloat float64
    var x MyFloat = 3.4

    // reflection object ? TYPE VALUE

    // Type is the representation of a Go type
    t := reflect.TypeOf(x)
    fmt.Println("Type:", t)
    fmt.Println("Type.String():", t.String())
    fmt.Println("Type.Kind():", t.Kind()) 

    // Value is the reflection interface to a Go value
    v := reflect.ValueOf(x)
    fmt.Println("Value:", v)
    fmt.Println("Value.String():", v.String()) // Value is interface type => <main.MyFloat Value>
    fmt.Println("Value.Type():", v.Type())  // Value.Type() => value's static type => main.MyFloat
    fmt.Println("Value.Kind():", v.Kind())  // Value.Kind() => value's underlying type => float64
    fmt.Println("Value.Float():", v.Float())

    y := v.Interface().(MyFloat) // y will have type float64.
    fmt.Println(y)
    fmt.Println("Value.Interface():", v.Interface())
    fmt.Printf("value is %7.1e\n", v.Interface())

    fmt.Println("settability of v:", v.CanSet())
    // v.SetFloat(7.1) // Error: will panic: reflect.Value.SetFloat using unaddressable value

    p := reflect.ValueOf(&x) // Note: take the address of x.
    fmt.Println("type of p:", p.Type())
    fmt.Println("settability of p:", p.CanSet())
    v2 := p.Elem()
    fmt.Println("settability of v2:", v2.CanSet())
    v2.SetFloat(7.1)
    fmt.Println(v2.Interface())
    fmt.Println(x)

    var w io.Writer = os.Stdout
    fmt.Println(reflect.TypeOf(w)) // "*os.File"

    StructReflect()
}
