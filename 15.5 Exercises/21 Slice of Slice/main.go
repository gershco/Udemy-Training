package main

import (
	"fmt"
)

func main() {

	sequence := []int{2, 3, 5, 8, 13, 21}

	for i := 0; i < len(sequence); i++ {

		fmt.Println("index number", i, "has the value", sequence[i])
	}

	fmt.Println()

	sliceofsequence := sequence[1:4]

	for i := 0; i < len(sliceofsequence); i++ {

		fmt.Println("Sliceofsequence index number", i, "has the value", sliceofsequence[i])
	}
}
