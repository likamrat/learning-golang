package main

import (
	"fmt"
)

func main() {
	favSport := "Basketball"
	switch favSport {
	case "Soccer":
		fmt.Println("Is this your favorite sport?")
	case "Basketball":
		fmt.Println(favSport, "is my favorite sport")
	}
}
