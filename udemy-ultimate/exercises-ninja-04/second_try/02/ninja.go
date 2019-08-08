package main

import "fmt"

func main() {
	x := []int{22, 6, 33, 11, 58, 23, 12, 68, 191, 2}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
