package main

import "fmt"

func main() {
	favSport := "Basketball"
	switch favSport {
	case "Soccer":
		fmt.Println("NOPE!")
	case "Tennis":
		fmt.Println("NOPE!")
	case "Basketball":
		fmt.Println(favSport, "is my favorite sport!")
	}
}
