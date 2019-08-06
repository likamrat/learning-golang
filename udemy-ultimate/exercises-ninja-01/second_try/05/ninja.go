package main

import (
	"fmt"
)

type gizmo int

var x gizmo
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(x)
	fmt.Printf("%T\n", y)
}
