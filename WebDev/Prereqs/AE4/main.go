package main

import "fmt"

type person struct {
	fName, lName string
}

func main() {

	p1 := person{"Fred", "Bloggs"}

	fmt.Println(p1)
	fmt.Printf(p1.lName)
}
