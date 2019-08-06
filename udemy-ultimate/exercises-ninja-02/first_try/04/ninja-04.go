package main

import (
	"fmt"
)

func main() {
	y := 4
	fmt.Printf("%d\t\t%b\t\t%#x\n", y, y, y)

	z := y << 1
	fmt.Printf("%d\t\t%b\t\t%#x", z, z, z)
}
