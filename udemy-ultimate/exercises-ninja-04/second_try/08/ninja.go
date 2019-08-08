package main

import "fmt"

func main() {
	pepole := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, v := range pepole {
		fmt.Println("The key is:", k, "\nThe values are:\n", v)
	}
}
