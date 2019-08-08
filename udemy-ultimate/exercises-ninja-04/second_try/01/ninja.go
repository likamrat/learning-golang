package main

import "fmt"

func main() {
	x := [5]int{5, 6, 8, 33, 21}

	for i, v := range x {
		fmt.Println("The index is", i, "and the value is", v)
	}
	fmt.Printf("%T\t", x)
}
