package main

import (
	"fmt"
)

func main() {
	x := 70
	if x == 40 {
		fmt.Println("our value is 40")
	} else if x == 41 {
		fmt.Println("our value is 41")
	} else if x == 42 {
		fmt.Println("our value is 42")
	} else if x == 43 {
		fmt.Println("our value is 43")
	} else {
		fmt.Println("our value is not 40, 41, 42 or 43")
		fmt.Println("our value is", x)
	}
}
