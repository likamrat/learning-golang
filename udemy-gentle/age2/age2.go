// example if statement with else if

package main

import (
	"fmt"
)

func main() {
	age := 1

	if age < 13 {
		fmt.Println("Wow, you're young!")
	} else if age < 20 {
		fmt.Println("You are a teenager")
	} else if age < 30 {
		fmt.Println("You are in your twenties")
	} else if age < 40 {
		fmt.Println("You are in your thirties")
	} else if age < 50 {
		fmt.Println("You are getting there!")
	} else if age >= 50 {
		fmt.Println("Over the hill!")
	}

}
