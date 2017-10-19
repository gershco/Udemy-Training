package main

import (
	"fmt"
)

func printarray(c [3][2]string) {
	for _, v1 := range c {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}

	var b [3][2]string
	b[0][0] = "Apple"
	b[0][1] = "Samsung"
	b[1][0] = "Microsoft"
	b[1][1] = "Google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"

	printarray(a)
	fmt.Printf("\n")
	printarray(b)
}
