package main

import "fmt"

func main() {
	// MAPS are KEY/VALUE store.
	m := map[string]int{ // This is our composite literal TYPE. It's a long TYPE but still a TYPE!
		"Lior":  35, // "Lior" is the KEY and "35" is the VALUE
		"Irena": 34,
		"Meir":  73,
	}
	fmt.Println(m)
	fmt.Println(m["Lior"])

	fmt.Println(m["Revital"]) // There is no Key/Value for "Revital" in the map so the value will be 0

	v, ok := m["Revital"] // This is a way to check if the Key/Value exist in the map
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Revital"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v) // This is false so it doesn't run
	}
	if v, ok := m["Irena"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v) // This is true so it will run
	}
}
