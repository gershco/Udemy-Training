package main

import (
	"fmt"
)

func main() {

	f := 25

	fmt.Print("\nf is of type ")

	fmt.Printf("%T", f)

	fmt.Print("\n\nfloat64(f) is of type ")

	fmt.Printf("%T\n", float64(f))

}
