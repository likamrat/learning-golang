package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person      // a secret agent is everything a person is + a license to kill
	ltk    bool // ltk = license to kill
}

func main() {
	secretAgent1 := secretAgent{
		person: person{
			first: "Lior",
			last:  "Kamrat",
			age:   35,
		},
		ltk: true,
	}

	person2 := person{
		first: "Irena",
		last:  "Kamrat",
		age:   34,
	}

	fmt.Println(secretAgent1)
	fmt.Println(person2)
	fmt.Println(secretAgent1.first, secretAgent1.last, secretAgent1.age, secretAgent1.ltk)
	fmt.Println(person2.first, person2.last, person2.age)
}
