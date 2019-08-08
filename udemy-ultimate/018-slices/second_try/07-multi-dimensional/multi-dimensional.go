package main

import "fmt"

func main() {
	marvel := []string{"Spiderman", "Deadpool", "Wolverine"}
	fmt.Println(marvel)
	dc := []string{"Superman", "Batman", "Flash"}
	fmt.Println(dc)

	xp := [][]string{marvel, dc}
	fmt.Println(xp)
}
