package main

import "fmt"

func main() {
	player := "LeBron"
	switch player {
	case "Kobe", "Jordan", "LeBron":
		fmt.Println("One of these 3 legends")
	case "NBA":
		fmt.Println("The game")
	default:
		fmt.Println("The default case")
	}
}
