package main

import (
	"fmt"
)

func main() {

	var e interface{}

	fmt.Printf("%v, %T\n\n", e, e)

	e = 15

	fmt.Printf("%v, %T\n\n", e, e)

	e = "frog"

	fmt.Printf("%v, %T\n\n", e, e)

	e = 24.675

	fmt.Printf("%v, %T\n\n", e, e)

}
