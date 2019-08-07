package main

import (
	"fmt"
)

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true) // or
	fmt.Println(true || false)
	fmt.Println(!true)
}
