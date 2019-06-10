package main

import "fmt"

func main() {
	movies := []string{"Forest Gump", "The Lior King", "Aladin"}
	tv := []string{"GOT", "Lost", "Breaking Bad"}
	fmt.Println(movies)
	fmt.Println(tv)

	favorite := [][]string{movies, tv} // this the new type. a SLICE of a SLICE of string [][]string
	fmt.Println(favorite)

	for i, xs := range favorite {
		fmt.Println("record: ", i)
		for j, value := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, value)
		}
	}
}
