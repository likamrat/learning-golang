package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 { // any number that is divided by 3 that equals 0 - print!
			fmt.Println(i)
		}
	}
}
