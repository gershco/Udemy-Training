package main

import "fmt"

func main() {

	a := [...]float64{67.7, 89.8, 21, 78}

	for i, v := range a { //range returns both the index and value

		fmt.Printf("The %d element of a is %.1f\n", i, v)

	}

}
