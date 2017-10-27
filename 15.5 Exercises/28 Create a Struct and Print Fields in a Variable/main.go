package main

import "fmt"

type person struct {
	firstName, lastName, hair string
	age                       int
}

func main() {

	per1 := person{
		firstName: "Fred",
		lastName:  "Bloggs",
		hair:      "blond",
		age:       24,
	}

	per2 := person{
		firstName: "John",
		lastName:  "Smith",
		hair:      "brown",
		age:       56,
	}

	per3 := person{"Frank", "Bolton", "grey", 67}

	fmt.Println("Person 1:", per1)
	fmt.Println()

	fmt.Println("Person 2:", per2)
	fmt.Println()

	fmt.Println("Person 3:", per3)
	fmt.Println()
}
