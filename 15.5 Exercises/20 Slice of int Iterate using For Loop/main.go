package main

import (
	"fmt"
)

func main() {

	myslice := []int{2, 3, 5, 8, 13, 21}

	for i := 0; i < len(myslice); i++ {

		fmt.Println("index number", i, "has the value", myslice[i])
	}
}
