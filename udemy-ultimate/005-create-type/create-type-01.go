package main

import (
	"fmt"
)

var y int

type hotdog int // We are creating a new type named "hotdog" and the underline type is int

var z hotdog

func main() {
	y = 100
	z = 200
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	// y = z 				// You cannot do this! "y" is of type int and "z" is of type "hotdog"
	// fmt.Println(y, z)
}
