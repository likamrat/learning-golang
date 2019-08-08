package main

import (
	"fmt"
)

func main() {
	usStates := make([]string, 50, 500) // 50 = Slice size / 500 = Array size

	// This solution is with composite literal
	usStates = []string{" Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut", " Delaware", " Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas", " Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts", " Michigan", " Minnesota", " Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey", " New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon", " Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", " Tennessee", " Texas", " Utah", " Vermont", " Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming"}
	fmt.Println(usStates)
	fmt.Println(len(usStates))
	fmt.Println(cap(usStates))

	for i := 0; i < len(usStates); i++ {
		fmt.Println(i, usStates[i])
	}
}
