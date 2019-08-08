package main

import "fmt"

func main() {
	x := []int{4, 88, 87, 99, 3}
	fmt.Println(x)
	fmt.Println(x[1])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println("")

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
