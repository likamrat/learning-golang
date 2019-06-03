package main

import (
	"fmt"
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x) // %d for decimal, %b for binary

	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)
}
