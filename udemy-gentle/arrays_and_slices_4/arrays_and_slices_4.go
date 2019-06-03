// arrays and slices example

package main

import "fmt"

var gGroceries []string // slice of strings

func addGrocery(a string) {
	fmt.Println("Capacity:", cap(gGroceries))
	gGroceries = append(gGroceries, a)
}

/*func listGroceries() {
	fmt.Println("Groceries list is as follows:")
	for i := 0; i < len(gGroceries); i++ {
		fmt.Println(gGroceries[i])
	}
}*/

/*func listGroceries() {
	for i, d := range gGroceries {
		fmt.Println(i, d)
	}
}*/

func listGroceries() {
	for _, d := range gGroceries { //removing the index (i) using _
		fmt.Println(d)
	}
}

func main() {
	addGrocery("Milk")
	addGrocery("Coffee")
	addGrocery("Banana")
	addGrocery("Apples")
	listGroceries()
	addGrocery("Celery")
	addGrocery("Newspapper")
	addGrocery("Pinapple")
	listGroceries()
}
