package main

import "fmt"

func main() {
	m := map[string]int{
		"MJ": 23,
		"KB": 8,
	}
	fmt.Println(m)
	fmt.Println("")

	m["KD"] = 35

	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println(m)
	fmt.Println("")

	// to use range over a slice
	xi := []int{66, 33, 88, 2, 9}
	for index, value := range xi {
		fmt.Println(index, value)
	}
}
