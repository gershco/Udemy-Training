package main

import "fmt"

func main() {

	prime := make([]int, 5, 10)

	fmt.Println()
	fmt.Println("The slice is made with a length of 5 and a capacity of 10: ", prime)

	prime[0] = 2
	prime[1] = 3
	prime[2] = 5
	prime[3] = 7
	prime[4] = 11

	fmt.Println()
	fmt.Println("Five integer values are assigned to the slice: ", prime)

	fmt.Printf("Length of prime is %d. Capacity of prime is %d.", len(prime), cap(prime))
	fmt.Println()

	//Append two new integers

	prime = append(prime, 13, 17, 19, 23)

	fmt.Println()
	fmt.Println("After adding four more integers, this is the slice: ", prime)

	fmt.Printf("Length of prime is %d. Capacity of prime is %d.", len(prime), cap(prime))
	fmt.Println()

	//Append four more integers

	prime = append(prime, 29, 31, 37)

	fmt.Println()
	fmt.Println("After adding three more integers, this is the slice: ", prime)

	fmt.Printf("Length of prime is %d. Capacity of prime is %d.\n\n", len(prime), cap(prime))

	fmt.Println("Note that the capacity has doubled.")
}
