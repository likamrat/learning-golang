package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello Go!")
	fmt.Println(n)
	fmt.Println(err)
	foo()
	g, _ := fmt.Println("Hello Go!")
	fmt.Println(g)
	foo()
}

func foo() {
	fmt.Println("I'm in foo!")
}
