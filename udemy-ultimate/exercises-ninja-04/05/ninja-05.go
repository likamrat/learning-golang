package main

import "fmt"

func main() {
	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range mySlice {
		fmt.Println(i, v)
	}
	mySlice = append(mySlice[:3], mySlice[6:]...)
	fmt.Println(mySlice)
}
