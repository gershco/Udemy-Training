package main

import ("fmt"
	"math")

func main() {
	
	fmt.Println("Here are the first 10 square numbers.")

	for i := 1; i <= 10; i++ {
		
		fmt.Println(sq(i))

	}}	

	// Creates a function to square a given integer
func sq (e int) float64 {
	g := float64(e)
		return math.Pow(g,2)
	}	