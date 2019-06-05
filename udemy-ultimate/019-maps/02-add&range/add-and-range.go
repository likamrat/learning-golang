package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Lior":  35,
		"Irena": 34,
		"Meir":  73,
	}
	fmt.Println(m["Lior"])

	// add a value to a map
	m["Eyal"] = 50
	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key, value)
	}

	xi := []int{77, 9, 5, 3, 99}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
