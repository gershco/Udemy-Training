package main

import (
	"fmt"
)

func main() {

	for no, i := 10, 1; no <= 19 && i <= 10; i, no = i+1, no+1 { //multiple initialisation and increment

		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

}
