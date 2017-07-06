package main

import (
  "fmt"
  "math"
)

/*
Outside a function, every statement begins with a keyword: 
var
func
const
type
*/


/*
Go's basic types are:

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

var b bool
const World = "世界"
var c, python, java bool
var i = 1 // If an initializer is present, the type can be omitted; the variable will take the type of the initializer
//j := 2 // error. Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

func main()  {
  var i int // A var statement can be at package or function level
  fmt.Println(c, python, java, i)

  // Short variable declarations
  k := 3 // Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type
  const Pi = 3.14 // Constants cannot be declared using the := syntax

	fmt.Printf("type:%T value:%v\n", b, b)
  fmt.Printf("type:%T value:%v\n", k, k)
  fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

  // Variables declared without an explicit initial value are given their zero value
  var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

  // Type conversions
  var x, y int = 3, 4
  f = math.Sqrt(float64(x*x + y*y)) // Go assignment between items of different type requires an explicit conversion
  // f = math.Sqrt(k) // error: cannot use k (type int) as type float64 in argument to math.Sqrt
  z := uint(f)
  fmt.Println(x, y, f, z)
}
