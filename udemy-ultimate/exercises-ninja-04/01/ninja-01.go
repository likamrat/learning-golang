package main

import "fmt"

func main() {
	myArray := [5]int{77, 44, 33, 55, 66}
	fmt.Println(myArray)
	myArray[0] = 4
	myArray[1] = 7
	myArray[2] = 9
	myArray[3] = 11
	myArray[4] = 2
	fmt.Println(myArray)
	fmt.Println(len(myArray))
	for i, v := range myArray {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\t", myArray)
}
