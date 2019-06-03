// variadic functions example

package main

import "fmt"

var gGroceries []string // slice of strings

func addGrocery(a ...string) {
	for _, d := range a {
		gGroceries = append(gGroceries, d)
	}
}

func listGroceries() {
	for _, d := range gGroceries {
		fmt.Println(d)
	}
}

func main() {
	addGrocery("Milk", "Coffee", "Banana")
	addGrocery("Apples")
	addGrocery("Celery")
	addGrocery("Newspapper")
	addGrocery("Pinapple")
	listGroceries()
}
