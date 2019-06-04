package main

import (
	"fmt"
)

func main() {
	x := []int{55, 77, 66, 23, 101} // the SLICE x
	fmt.Println(x)
	for i, v := range x { // I am just practicing on range
		fmt.Println(i, v)
	}
	fmt.Println("")
	x = append(x, 99, 4, 300, 31)
	fmt.Println(x)
	for i, v := range x { // I am just practicing on range again
		fmt.Println(i, v)
	}
	fmt.Println("")
	// APPEND a SLICE to a SLICE
	y := []int{95, 1092, 521}
	x = append(x, y...) // the "..." means take all the values in a data structure (the SLICE of "y") and add them to the end.
	fmt.Println(x)
	for i, v := range x { // I am just practicing on range again
		fmt.Println(i, v)
	}
	fmt.Println("")
	fmt.Println(x[8]) // just random practice

}
