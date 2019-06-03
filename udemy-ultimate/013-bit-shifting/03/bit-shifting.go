package main

import (
	"fmt"
)

const (
	kb = 1024
	mb = 1024 * kb
	gb = 1024 * mb
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b", gb, gb)
}
