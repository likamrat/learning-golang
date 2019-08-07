package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("this should print")
		fallthrough
	case (6 == 6):
		fmt.Println("this should also print")
		fallthrough
	case (7 == 9):
		fmt.Println("why are you printing this if it's false?!")
		fallthrough
	case (6 == 9):
		fmt.Println("why are you printing this if it's false?!")
	}
}
