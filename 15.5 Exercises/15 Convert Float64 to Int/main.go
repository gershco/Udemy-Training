package main

import (
	"fmt"
)

func main() {

	var f float64

	f = 2.6

	fmt.Print("\nf, of type ")

	fmt.Printf("%T,", f)

	fmt.Print(" is assigned the value ")

	fmt.Print((f), ".")

	fmt.Print("\n\nint(f), of type ")

	fmt.Printf("%T,", int(f))

	fmt.Print(" has the value ")

	fmt.Print(int(f), ".\n")

}
