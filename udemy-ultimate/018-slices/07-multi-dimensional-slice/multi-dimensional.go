package main

import (
	"fmt"
)

func main() {
	family := []string{"Irena", "Revital", "Mom", "Dad"}
	friends := []string{"Yaron", "Catalin", "Tomer"}
	fmt.Println(family)
	fmt.Println(friends)

	xp := [][]string{family, friends}
	fmt.Println(xp)
}
