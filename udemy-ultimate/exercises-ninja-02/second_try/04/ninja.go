package main

import "fmt"

func main() {
	a := 4
	fmt.Printf("%d\t\t%b\t\t%#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\t\t%b\t\t%#x\n", b, b, b)
}
