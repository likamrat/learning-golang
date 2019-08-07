package main

import "fmt"

func main() {
	x := 89
	if x == 23 {
		fmt.Println("our value is 23")
	} else if x == 35 {
		fmt.Println("our value is", x)
	} else if x == 66 {
		fmt.Println("our value is", x)
	} else {
		fmt.Println("our value is not 23, 35 or 66, it is", x)
	}
}
