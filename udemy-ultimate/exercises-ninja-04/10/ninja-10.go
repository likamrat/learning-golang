package main

import "fmt"

func main() {
	myMap := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(myMap)
	fmt.Println("")

	myMap["Lior"] = []string{`Azure`, `golang`, `whatever`}

	delete(myMap, "no_dr")

	if value, ok := myMap["no_dr"]; ok {
		fmt.Println("This should not be printed", value) // change the value in the if statement to KEY you didn't delete to see if it works
	}

	for k, v := range myMap {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
