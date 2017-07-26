package main

import "fmt"
import "os"
import "io"
import "reflect"

// A type assertion provides access to an interface value's underlying concrete value. 
// t := i.(T)
// t, ok := i.(T)
// That type T must either be the concrete type held by the interface, or a second interface type that the value can be converted to

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
    //f1 := i.(float64) // panic
    f1, _ := i.(float64) // ignore panic error
    fmt.Println(f1, ok)

    var r io.Reader
    var w io.Writer
    tty, _ := os.OpenFile("/dev/tty", os.O_RDWR, 0)
    r = tty // r contains, schematically, the (value, type) pair, (tty, *os.File)
    w = r.(io.Writer)
    fmt.Println(reflect.TypeOf(r)) // the type *os.File implements methods Read 
    fmt.Println(reflect.TypeOf(r).Kind())
    fmt.Println(reflect.TypeOf(w)) // r also implements io.Writer
}
