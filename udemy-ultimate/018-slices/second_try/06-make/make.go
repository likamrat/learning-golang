package main

import "fmt"

func main() {
	// make(slice type, length, capacity)
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // 100 spots for the underlay array to use

	x[0] = 19
	x[9] = 55
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// x[10] = 102 // This will cause "Index out of range" since the length of the slice is 10 and [10] is index #11
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	x = append(x, 211)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 212)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 213) // This will make slice length go to 13 which will force the complier to delete the old underline array and create new array double the size (24)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
