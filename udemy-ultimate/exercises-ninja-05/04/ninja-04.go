package main

import "fmt"

func main() {
	person := struct {
		firstName string
		lastName  string
		over30    bool
		favMovies map[string]int
		favDrinks []string
	}{
		firstName: "Lior",
		lastName:  "Kamrat",
		over30:    true,
		favMovies: map[string]int{
			"Forest Gump":   1,
			"The Lion King": 2,
			"Die Hard":      3,
		},
		favDrinks: []string{
			"Beer",
			"Water",
			"Orange Juice",
		},
	}

	fmt.Println(person.firstName)
	fmt.Println(person.lastName)
	fmt.Println("Is this person over 30?", person.over30)
	fmt.Println(person.favMovies)
	fmt.Println(person.favDrinks)

	for key, value := range person.favMovies {
		fmt.Println(key, value)
	}

	for index, value := range person.favDrinks {
		fmt.Println(index, value)
	}

}
