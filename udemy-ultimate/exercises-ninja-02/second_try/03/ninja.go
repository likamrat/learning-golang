package main

import "fmt"

const (
	x         = 23
	y float32 = 24.5
)

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
