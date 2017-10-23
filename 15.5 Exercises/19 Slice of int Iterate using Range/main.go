package main

import (
	"fmt"
)

func main() {

	myslice := []int{2, 3, 4, 5, 6, 8}

	//range returns two values for each iteration - the index followed by a copy of the element at that index.

	for i, v := range myslice {

		fmt.Println("index number", i, "has the value", v)
	}
	fmt.Println()
	fmt.Println("Blank identifier was used in the for loop to remove the index.")

	for _, v := range myslice {

		fmt.Println(v)
	}

}
