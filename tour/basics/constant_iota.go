package main

import "fmt"

// iota: Elegant Constants in Golang

type Allergen int
const (
	IgEggs Allergen = 1 << iota  // 1 << 0 which is 00000001
	IgChocolate					 // 1 << 1 which is 00000010
	IgNuts						 // 1 << 2 which is 00000100
	IgStrawberries				 // 1 << 3 which is 00001000
	IgShellfish                  // 1 << 4 which is 00010000
)

type ByteSize float64
const (
    _           = iota                   // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota) 		 // 1 << (10*1)
    MB                                   // 1 << (10*2)
    GB                                   // 1 << (10*3)
    TB                                   // 1 << (10*4)
    PB                                   // 1 << (10*5)
    EB                                   // 1 << (10*6)
    ZB                                   // 1 << (10*7)
    YB                                   // 1 << (10*8)
)

func main() {
	fmt.Printf("%#08b\n", IgEggs | IgChocolate | IgShellfish)
	fmt.Printf("%f %f %f %f %f %f\n", KB, MB, GB, EB, ZB, YB)
}