package main

import (
	"fmt"
)

func main() {
	n := "Bond"
	switch n {
	case "Money":
		fmt.Println("no print")
	case "Bond":
		fmt.Println("Bond, James Bond")
	}
}
