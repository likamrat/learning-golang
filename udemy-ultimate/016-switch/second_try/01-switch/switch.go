package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print")
	case (3 == 3):
		fmt.Println("This should print")
	case (2 == 2):
		fmt.Println("This should print, does it print?")
	}
}
