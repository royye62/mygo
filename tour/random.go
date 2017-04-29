package main

import (
  "time"
	"fmt"
	"math/rand"
)

func main() {
  rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(1000))
}
