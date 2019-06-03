package main

import (
	"fmt"
)

func main() {
	x := 60
	if x == 42 {
		fmt.Println("our value is", x)
	} else if x == 43 {
		fmt.Println("our value is", x)
	} else {
		fmt.Println("It's not 42 or 43! our value is", x)
	}
}
