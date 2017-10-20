package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 2000; i++ {

		fmt.Printf("%d \t %b \t \t %X \t \t %U \t \t %q \n", i, i, i, i, i)
	}

}
