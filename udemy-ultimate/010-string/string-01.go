package main

import (
	"fmt"
)

func main() {
	s := "hello world!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)        // a slice of byte
	fmt.Println(bs)        // the output is in ascii
	fmt.Printf("%T\n", bs) // byte is an alias of unit8
}
