package main

import "fmt"

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {
	person1 := person{
		first: "Lior",
		last:  "Kamrat",
		favIceCream: []string{
			"Chocolate",
			"Vanilla",
			"Lemon",
		},
	}

	person2 := person{
		first: "Irena",
		last:  "Kamrat",
		favIceCream: []string{
			"Cookies",
			"Caramel",
			"Mango",
		},
	}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println("")

	for index, value := range person1.favIceCream {
		fmt.Println(index, value)
	}
	for index, value := range person2.favIceCream {
		fmt.Println(index, value)
	}
}
