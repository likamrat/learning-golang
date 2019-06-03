// arrays and slices example

package main

import "fmt"

const gCap int = 5          // total number of array elements
var gLength int             // integer number of items in our array currently
var gGroceries [gCap]string // array of strings

func addGrocery(a string) {
	if gLength < gCap {
		gGroceries[gLength] = a
		gLength++
	} else {
		fmt.Println("Error!!! We've reached the upper bound of the array!!!")
	}
}

func listGroceries() {
	fmt.Println("Groceries list is as follows:")
	for i := 0; i < gLength; i++ {
		fmt.Println(gGroceries[i])
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
	listGroceries()
}
