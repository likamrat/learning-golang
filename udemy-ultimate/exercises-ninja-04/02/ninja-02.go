package main

import "fmt"

func main() {
	mySlice := []int{99, 4, 100, 200, 66, 23, 9, 403, 44, 999}
	for i, v := range mySlice {
		fmt.Println("This the index:", i, "This is the value:", v)
	}
	fmt.Printf("%T\t", mySlice)
}
