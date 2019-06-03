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
	//y = z      // you can't do this! you can't use type hotdog and assign it to a VARIABLE that was DECLARED as type int.
	y = int(z) // Because we know the underline type of hotdog is int, we can use CONVERSION using T()
	fmt.Println("y =", y, "\nz =", z)
}
