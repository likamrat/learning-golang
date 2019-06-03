// example for switch

package main

import (
	"fmt"
)

func main() {

	for day := 1; day <= 12; day++ {
		fmt.Println("On the", day, "of Christmeas my true love sent to me:")
		switch day {
		case 12:
			fmt.Println("12 Drummers Drumming")
		case 11:
			fmt.Println("11 Pipers Piping")
		case 10:
			fmt.Println("10 lords a leaping ")
		case 9:
			fmt.Println("9 ladies dancing")
		case 8:
			fmt.Println("8 maids a milking")
		case 7:
			fmt.Println("7 swans a swiming")
		case 6:
			fmt.Println("6 geese a laying")
		case 5:
			fmt.Println("5 golden rings")
		case 4:
			fmt.Println("4 calling birds")
		case 3:
			fmt.Println("3 french hens")
		case 2:
			fmt.Println("2 turtle doves")
		case 1:
			fmt.Println("a Partridge in a Pear tree")
		}
	}
}
