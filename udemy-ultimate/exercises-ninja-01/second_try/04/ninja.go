package main

import (
		"fmt"
)

type gizmo int

var x gizmo

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
