package main

import (
	"fmt"
)

func main() {
	n := "Bond"
	switch n {
	case "Money", "Bond", "Dr. No":
		fmt.Println("Money or Bond or Dr. No")
	case "Q":
		fmt.Println("no print")
	}
}
