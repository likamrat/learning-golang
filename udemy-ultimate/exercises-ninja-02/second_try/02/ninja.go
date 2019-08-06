package main

import "fmt"

func main() {
	a := (42 == 23)
	b := (42 <= 23)
	c := (42 >= 23)
	d := (42 != 23)
	e := (42 < 23)
	f := (42 > 23)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
