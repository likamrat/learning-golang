package main

import (
	"fmt"
)

func main() {
	x := []int{23, 44, 76, 98, 6}
	fmt.Println(len(x)) // get the length of a slice
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x { // RANGE over a SLICE
		fmt.Println(i, v)
	}
}
