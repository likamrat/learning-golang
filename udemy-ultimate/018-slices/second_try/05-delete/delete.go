package main

import "fmt"

func main() {
	x := []int{6, 7, 33, 23, 5}
	fmt.Println(x)
	x = append(x, 11, 22, 4, 102)
	fmt.Println(x)
	fmt.Println("")

	y := []int{3, 5, 100}
	x = append(x, y...)
	fmt.Println(x)
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
