package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

// Multiple results
func swap(x, y string) (string, string) {
  return y, x
}

// Named return values
func split(sum int) (x, y int)  {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main()  {

  fmt.Println(add(42, 13))

  a, b := swap("hello", "world");
  fmt.Println(a, b)

  fmt.Println(split(17))
}
