package main

import (
	"fmt"
)

func main() {

	a := 394

	fmt.Print("The decimal number ", a, " is ")
	fmt.Printf("%b", a)
	fmt.Print(" in binary and ")
	fmt.Printf("%x", a)
	fmt.Print(" in hexadecimal.")

}
