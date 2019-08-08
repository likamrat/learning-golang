package main

import "fmt"

func main() {
	m := map[string]int{
		"Lior":    35,
		"Irena":   34,
		"Revital": 43,
	}
	fmt.Println(m)
	delete(m, "Revital")
	fmt.Println(m)

	if v, ok := m["Revital"]; ok { // we can use the ", ok" to check if a value was actually deleted from the map
		fmt.Println(v)
	}
	if v, ok := m["Lior"]; ok {
		fmt.Println(v)
	}
}
