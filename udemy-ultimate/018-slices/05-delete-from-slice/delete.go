package main

import (
	"fmt"
)

func main() {
	x := []int{77, 66, 88, 3, 7, 4}
	fmt.Println(x)
	y := []int{1000, 2000, 3000, 4000}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("")
	for i, v := range x {
		fmt.Println(i, v)
	}

	// DELETE from a SLICE
	// Let's say I want to delete the value 4 & 1000
	x = append(x[:5], x[7:]...)
	// there is no value BEFORE the ":" in the slice x[:5] because we are sliceing from the begining
	// there is no value AFTER the ":" in the slice x[7:] because we are sliceing all the way to the end
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}
}
