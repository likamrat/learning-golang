// example for switch

package main

import (
	"fmt"
)

func main() {

	fmt.Println("On the", 1, "of Christmeas my true love sent to me:")
	fmt.Println("a Partridge in a Pear tree")
	fmt.Println("")

	for day := 2; day <= 12; day++ {
		fmt.Println("On the", day, "of Christmeas my true love sent to me:")
		switch day {
		case 12:
			fmt.Println("12 Drummers Drumming")
			fallthrough
		case 11:
			fmt.Println("11 Pipers Piping")
			fallthrough
		case 10:
			fmt.Println("10 lords a leaping ")
			fallthrough
		case 9:
			fmt.Println("9 ladies dancing")
			fallthrough
		case 8:
			fmt.Println("8 maids a milking")
			fallthrough
		case 7:
			fmt.Println("7 swans a swiming")
			fallthrough
		case 6:
			fmt.Println("6 geese a laying")
			fallthrough
		case 5:
			fmt.Println("5 golden rings")
			fallthrough
		case 4:
			fmt.Println("4 calling birds")
			fallthrough
		case 3:
			fmt.Println("3 french hens")
			fallthrough
		case 2:
			fmt.Println("2 turtle doves")
			fallthrough
		case 1:
			fmt.Println("and a Partridge in a Pear tree")
		}

		fmt.Println("")
	}
}
