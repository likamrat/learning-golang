package main

import "fmt"

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	xi := [][]string{xs1, xs2}
	fmt.Println(xi)
	for sliceIndex, xs := range xi {
		fmt.Println("The index for the slice is", sliceIndex)
		for sliceSliceIndex, value := range xs {
			fmt.Println("\tIndex position:\t", sliceSliceIndex, "\tValue:\t", value)
		}
	}
}
