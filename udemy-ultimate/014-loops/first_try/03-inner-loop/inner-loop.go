package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ { // init; condition; post and then {code block} // this is the for loop structure
		fmt.Printf("This is the outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t This is the inner loop: %d\n", j)
		}
	}
}
