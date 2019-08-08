package main

import "fmt"

func main() {
	m := map[string]int{
		"MJ": 23,
		"KB": 8,
	}
	fmt.Println(m)
	fmt.Println(m["MJ"])
	fmt.Println(m["LeBron"]) // if you have a key that is not part of the map it will return the 0 value by default
	fmt.Println("")

	// we can check this using the ", ok"
	v, ok := m["LeBron"]
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println("")

	if v, ok := m["LeBron"]; ok {
		fmt.Println("This will not print", v)
	}
	fmt.Println("")
	if v, ok := m["KB"]; ok {
		fmt.Println("This will print", v)
	}
}
