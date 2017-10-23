package main

import "fmt"

func main() {

	fours := []int{4, 8, 12, 16, 20, 24, 28, 32}

	fmt.Println()
	fmt.Println(fours)

	fmt.Printf("Length of fours is %d. Capacity of fours is %d.", len(fours), cap(fours))

	fours = append(fours[1:2], fours[3:4], fours[5:6], fours[7:])
	fmt.Println(fours)
}
