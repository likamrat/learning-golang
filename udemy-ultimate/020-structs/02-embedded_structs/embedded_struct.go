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
	person1 := person{
		first: "Lior",
		last:  "Kamrat",
		age:   35,
	}

	person2 := person{
		first: "Irena",
		last:  "Kamrat",
		age:   34,
	}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.first, person1.last, person1.age)
	fmt.Println(person2.first, person2.last, person2.age)
}
