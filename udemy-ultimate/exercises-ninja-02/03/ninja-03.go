package main

import (
	"fmt"
)

const (
	a int = 5
	b     = 71.5
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
