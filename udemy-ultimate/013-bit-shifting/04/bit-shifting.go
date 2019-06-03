package main

import (
	"fmt"
)

const (
	_ = iota // We are throwing the first iota away using the "_". This will be the "zero iota"
	//kb = 1024
	kb = 1 << (iota * 10) // this will bit shift the 1 by 10 so 10000000000
	mb = 1 << (iota * 10) // this will bit shift the 1 by 10 so 100000000000000000000
	gb = 1 << (iota * 10) // this will bit shift the 1 by 10 so 1000000000000000000000000000000
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b", gb, gb)
}
