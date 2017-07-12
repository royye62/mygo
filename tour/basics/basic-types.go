package main


/* 
basic types:
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32, represents a Unicode code point
float32 float64
complex64 complex128

compound types:
array
struct

reference types:
slice
map
func
channel

interface type: represent fixed sets of methods, An interface variable can store any concrete (non-interface) value as long as that value implements the interface's methods.

*/

import (
    "fmt"
    "math"
    "math/cmplx"
)

var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)

    // zero value
    // 0 for numberic types
    // false for the boolean type
    // "" for the strings 
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)

    // Type conversions
    // The expression T(v) converts the value v to the type T
    // Unlike in C, in Go assignment between items of different type requires an explicit conversion.
    var x, y int = 3, 4
    var f2 float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f2)
    fmt.Println(x, y, z, f2)

    // Type inference
    // when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant
    v := 2 + 1 // int
    fmt.Printf("v is of type %T\n", v)
    v2 := 2 + 0.1 // float64
    fmt.Printf("v is of type %T\n", v2)
    v3 := 0.867 + 0.5i // complex128
    fmt.Printf("v is of type %T\n", v3)
}
