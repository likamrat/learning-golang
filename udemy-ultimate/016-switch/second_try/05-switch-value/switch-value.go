package main

import "fmt"

func main() {
	switch "Jordan" {
	case "Kobe":
		fmt.Println("Kobe Bryant")
	case "Jordan":
		fmt.Println("MJ")
	case "LeBron":
		fmt.Println("The King")
	}
	name := "LeBron"
	switch name {
	case "Kobe":
		fmt.Println("Kobe Bryant")
	case "Jordan":
		fmt.Println("MJ")
	case "LeBron":
		fmt.Println("The King")
	}
}
