package main

import "fmt"

type person struct {
	fName, lName string
	favFood      []string
}

func main() {

	p1 := person{"Fred", "Bloggs", []string{"chips", "pizza", "ice-cream"}}

	fmt.Println(p1.favFood)
	fmt.Println()

	for _, v := range p1.favFood {

		fmt.Printf("%s\n", v)
	}
}
