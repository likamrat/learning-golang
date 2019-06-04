package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 12) // the MAKE built-in function. Can save processing power so a new array will not have to be created everytime.
	fmt.Println(x)           // the SLICE
	fmt.Println(len(x))      // the SLICE length
	fmt.Println(cap(x))      // the UNDERLINE ARRAY capacity
	fmt.Println("")

	x = append(x, 99)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("")

	x[3] = 10
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("")

	y := []int{55, 6, 7}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // Since the UNDERLINE ARRAY capacity needs to be increased, the Go runtime will double (12 -> 24) the size of the UNDERLINE ARRAY
	fmt.Println("")
	for i, v := range x {
		fmt.Println(i, v)
	}
}
