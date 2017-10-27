package main

import (
	"fmt"
)

func subtractOne(numbers []int) {
	for i := range numbers {
		numbers[i]--
	}
}
func subtractTwo(numbers [3]int) {
	for i := range numbers {
		numbers[i] -= 2

	}
	fmt.Println()
	fmt.Println("The array after but within the function is ", numbers)
	fmt.Println()
}
func main() {

	evs := []int{10, 8, 6}
	fmt.Println("slice before function call is", evs)
	subtractOne(evs)                                 //function modifies the slice
	fmt.Println("slice after function call is", evs) //modifications are visible outside

	fmt.Println()

	odds := [3]int{9, 7, 5}
	fmt.Println("Array before function call is", odds)
	subtractTwo(odds)
	fmt.Println("Array after function call is", odds)

}
