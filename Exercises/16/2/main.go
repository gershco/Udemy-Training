package main

import "fmt"

func main() {

	var x = 18

	mine := func(n int) (int, bool) {

		return n / 2, n%2 == 0
	}
	fmt.Println(mine(x))
}
