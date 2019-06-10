package main

import (
	"fmt"
)

func main() {
	x := []int{56, 99, 8, 0, 44, 6}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[4])
	fmt.Println(x[1:4]) // In this case, the slice of a slice will not incl. the number in index position 4. "Accessing array by index postion"

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("")
	// ALTERNATIVE WAY:

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i]) // give me the index and then give me the VALUE from that index postion in the SLICE
	}

}
