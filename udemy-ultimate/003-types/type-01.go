package main

import (
	"fmt"
)

var y = 100

var z string = "Great success!"

// DECLARED the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "Great success!"
// Go is a STATIC programming language not DYNAMIC
// a VARIABLE is IDENTIFIER to hold a value of a certain TYPE

var a string = `Borat said, "Great success!"` // USE the `` in order to keep the ""

func main() {
	fmt.Println("y =", y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
