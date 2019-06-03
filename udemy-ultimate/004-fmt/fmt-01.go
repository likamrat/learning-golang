package main

import (
	"fmt"
)

var y = 100

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	fmt.Printf("%#x\n%b\n%x\n", y, y, y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)
	fmt.Printf("%#x\v%b\v%x\v", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
}
