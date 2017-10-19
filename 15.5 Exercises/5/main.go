package main

import (
	"fmt"
)

func main() {

	var a rune

	a = 23

	fmt.Printf("%T", a)

	// rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values.
}
