package main

import "fmt"

type person struct {
	fName, lName string
	favFood      []string
}

func (p person) walk() {
	fmt.Printf("%s is walking.", p.fName)
}

func main() {

	p1 := person{"Fred", "Bloggs", []string{"chips", "pizza", "ice-cream"}}

	p1.walk()
}
