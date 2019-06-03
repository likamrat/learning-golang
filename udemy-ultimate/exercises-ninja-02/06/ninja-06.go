package main

import (
	"fmt"
)

const (
	a = 2015 + iota
	year1
	year2
	year3
	year4
)

const (
	b      = 2015 + iota
	year10 = 2015 + iota
	year20 = 2015 + iota
	year30 = 2015 + iota
	year40 = 2015 + iota
)

func main() {
	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)

	fmt.Println("")

	fmt.Println(year10)
	fmt.Println(year20)
	fmt.Println(year30)
	fmt.Println(year40)
}
