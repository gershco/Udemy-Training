package main

import (
	"fmt"
)

func main() {

	s := []int{2, 4, 6, 8, 9}

	fmt.Println(s)

	for i := range s {
		fmt.Printf("%d.\n", i)
	}

	fmt.Println()

	for _, v := range s {
		fmt.Printf("%d\n", v)
	}

	fmt.Println()

	for i, v := range s {
		fmt.Printf("The %d element of s is %d.\n", i, v)
	}
}
