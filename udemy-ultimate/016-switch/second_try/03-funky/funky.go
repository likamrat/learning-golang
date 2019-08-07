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
		fallthrough
	case (2 == 2):
		fmt.Println("This should print, does it print?")
		fallthrough
	case (7 == 2):
		fmt.Println("This is not true")
		fallthrough
	case (4 == 4):
		fmt.Println("This is true")
		fallthrough // if the case is true (4 == 4), fallthrough to the next case
	default:
		fmt.Println("this is the default case")
	}

}
