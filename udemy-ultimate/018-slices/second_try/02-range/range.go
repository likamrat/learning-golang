package main

import "fmt"

func main() {
	x := []int{23, 44, 66, 77, 2}
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println("")
	for i, v := range x {
		fmt.Println(i, v)
	}
}
