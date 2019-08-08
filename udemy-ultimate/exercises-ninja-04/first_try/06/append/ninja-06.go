package main

import "fmt"

func main() {
	usStates := make([]string, 50, 500)
	fmt.Println(usStates)
	fmt.Println(len(usStates))
	fmt.Println(cap(usStates))

	// put a value into each of the 50 positions in the length of the slice
	// we are putting values into the 50 positions that are the length of the slice
	for i := 0; i < 50; i++ {
		usStates[i] = fmt.Sprintf("Position %d", i)
	}
	fmt.Println(usStates)
	fmt.Println(len(usStates))
	fmt.Println(cap(usStates))

	fmt.Println("")

	// append values to the slice grows the length of the slice
	// the underlying array "capacity" of 500 is the same
	usStates = append(usStates, " Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut", " Delaware", " Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas", " Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts", " Michigan", " Minnesota", " Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey", " New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon", " Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", " Tennessee", " Texas", " Utah", " Vermont", " Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming")
	fmt.Println(usStates)
	fmt.Println(len(usStates))
	fmt.Println(cap(usStates))
	for i := 0; i < len(usStates); i++ {
		fmt.Println(i, usStates[i])
	}
}
