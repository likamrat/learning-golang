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
		first: "James",
		last:  "Bond",
		favIceCream: []string{
			"Cookies",
			"Caramel",
			"Mango",
		},
	}

	myMap := map[string]person{
		person1.last: person1,
		person2.last: person2,
	}

	fmt.Println(myMap)

	for _, value := range myMap {
		// fmt.Println(key)
		fmt.Println(value.first)
		fmt.Println(value.last)
		for index, value1 := range value.favIceCream {
			fmt.Println(index, value1)
		}
		fmt.Println("----------")
	}

}
