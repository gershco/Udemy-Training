package main

import (
	"fmt"
)

func main() {
	veggies := []string{"carrot", "tomato", "pepper"}
	fruits := []string{"orange", "apple"}
	vegfr := append(veggies, fruits...)
	frveg := append(fruits, veggies...)

	fmt.Println()
	fmt.Println(`Slice of string "veggies" is`, veggies)
	fmt.Printf("It has a length of %d and a capacity of %d.\n", len(veggies), cap(veggies))
	fmt.Println()

	fmt.Println(`Slice of string "fruits" is`, fruits)
	fmt.Printf("It has a length of %d and a capacity of %d.\n", len(fruits), cap(fruits))
	fmt.Println()

	fmt.Println(`After running "vegfr := append(veggies, fruits...)", slice of string "vegfr" is`, vegfr)
	fmt.Printf("It has length %d and capacity %d.\n", len(vegfr), cap(vegfr))

	fmt.Println()

	fmt.Println(`After running "frveg := append(fruits, veggies...)", slice of string "frveg" is`, frveg)
	fmt.Printf("It has length %d and capacity %d.\n", len(frveg), cap(frveg))
	fmt.Println()
}

