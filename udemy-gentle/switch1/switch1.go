// example for switch

package main

import (
	"fmt"
)

func main() {
	for day := 1; day <= 12; day++ {
		if day == 12 {
			fmt.Println("12 Drummers Drumming")
		} else if day == 11 {
			fmt.Println("11 Pipers Piping")
		} else if day == 10 {
			fmt.Println("10 lords a leaping ")
		} else if day == 9 {
			fmt.Println("9 ladies dancing")
		} else if day == 8 {
			fmt.Println("8 maids a milking")
		} else if day == 7 {
			fmt.Println("7 swans a swiming")
		} else if day == 6 {
			fmt.Println("6 geese a laying")
		} else if day == 5 {
			fmt.Println("5 golden rings")
		} else if day == 4 {
			fmt.Println("4 calling birds")
		} else if day == 3 {
			fmt.Println("3 french hens")
		} else if day == 2 {
			fmt.Println("2 turtle doves")
		} else if day == 1 {
			fmt.Println("a Partridge in a Pear tree")
		}
	}
}
