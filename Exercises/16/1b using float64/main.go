package main

import "fmt"

func main() {
	var x = 25
	fmt.Println(half(x))
}

func half(n int) (float64, bool) {

	return float64(n) / 2, n%2 == 0
}
