package main

import "fmt"

func main() {

	prime := make([]int, 5, 10)

	fmt.Println(prime)

	prime[0] = 2
	prime[1] = 3
	prime[2] = 5
	prime[3] = 7
	prime[4] = 11

	fmt.Println()
	fmt.Println(prime)

	fmt.Println()
	fmt.Printf("Length of prime is %d. Capacity of prime is %d.", len(prime), cap(prime))
}
