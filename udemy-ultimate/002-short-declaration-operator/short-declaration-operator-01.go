package main

import (
	"fmt"
)

var y = 50 // var is on the package level

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE
	x := 100 // the best practice is to always use the short operators inside your function, Use it only where you need it.
	fmt.Println(x)
	x = 200
	fmt.Println(x)
	fmt.Println(y)

	foo()
}

func foo() {
	fmt.Println(y)
}
