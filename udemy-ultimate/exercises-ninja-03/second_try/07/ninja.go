package main

import "fmt"

func main() {
	x := 23
	if x == 100 {
		fmt.Println("x = to", x)
	} else {
		fmt.Println("x != from 100 and equal to", x)
	}
}
