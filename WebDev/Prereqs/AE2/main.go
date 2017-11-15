package main

import (
	"fmt"
)

func main() {

	grades := map[string]int{"Fred": 9, "Joe": 7, "Frank": 8, "Jon": 6, "Dan": 5}

	fmt.Println(grades)

	for key := range grades {
		fmt.Printf("%s\n", key)
	}

	fmt.Println()

	for _, value := range grades {
		fmt.Printf("%d\n", value)
	}

	fmt.Println()

	for key, value := range grades {
		fmt.Printf("%s\t %d\n", key, value)
	}

}
