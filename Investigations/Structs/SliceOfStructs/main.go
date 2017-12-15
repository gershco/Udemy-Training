package main

import "fmt"

//Declare structure type person.

type person struct {
	first, last string
	age         int
}

func main() {

	// Define p1 and p2 structures.

	p1 := person{"James", "Bond", 38}
	p2 := person{"Miss", "Moneypenny", 29}

	fmt.Println("Struct p1 is", p1)
	fmt.Println("Struct p2 is", p2)
	fmt.Println()

	// Create the slice of struct

	xp := []person{}

	// Add the p1 and p2 structs to the slice.

	xp = append(xp, p1)
	xp = append(xp, p2)

	fmt.Println("Slice of structs xp is", xp)
}
