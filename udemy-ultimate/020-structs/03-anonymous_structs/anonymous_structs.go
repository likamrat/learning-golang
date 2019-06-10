package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	person1 := struct { // this is an anonymous struct because it doesn't have a name
		first string
		last  string
		age   int
	}{
		first: "Lior",
		last:  "Kamrat",
		age:   35,
	}

	fmt.Println(person1)
	fmt.Println(person1.first, person1.last, person1.age)
}
