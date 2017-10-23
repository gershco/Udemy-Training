package main

import (
	"fmt"
)

func main() {

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}

	fmt.Printf("length of array%d capacity %d\n", len(fruitarray), cap(fruitarray))

	fruitslice := fruitarray[6:]

	for i, v := range fruitslice {

		fmt.Println("index number", i, "has the value", v)
	}

	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6

	fmt.Println()

	fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity

	for i, v := range fruitslice {

		fmt.Println("index number", i, "has the value", v)
	}

	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
}
