package main

import "fmt"

func main () {

	fmt.Println("Please enter two different numbers, the larger one first")

	var large int
	var small int

	fmt.Scan(&large, &small)

// Lines 16 and 18 could be ommitted, with large%small replacing rem in line 20.

	var rem int

	rem = large%small

	fmt.Println("When the calculation" , large , "÷" , small , "is performed, the remainder is" , rem)
}