package main

import "fmt"

func main() {
	m := map[string]int{
		"MJ": 23,
		"KB": 8,
	}
	fmt.Println(m)

	m["KD"] = 35

	fmt.Println(m)

	delete(m, "MJ")
	fmt.Println(m)

	if v, ok := m["KD"]; ok {
		fmt.Println("value", v)
		delete(m, "KD")
	}
	fmt.Println(m)
}
