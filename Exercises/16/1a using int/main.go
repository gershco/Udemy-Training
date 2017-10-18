package main

import "fmt"

func main() {
	var x = 18
	fmt.Println(half(x))
}

func half(n int) (int, bool) {

	return n / 2, n%2 == 0
}
