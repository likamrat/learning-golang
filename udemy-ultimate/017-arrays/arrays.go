package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 42 // zero-based index (start the count from 0)
	fmt.Println(x)
	fmt.Println(len(x)) // array length
}
