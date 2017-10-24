package main

import "fmt"

func main() {

	pythons := []string{"ball", "royal", "reticulated", "indian", "blood", "carpet"}

	fmt.Println()
	fmt.Println(pythons)

	fmt.Printf("Length of pythons is %d. Capacity of pythons is %d.\n", len(pythons), cap(pythons))

	boas := []string{"sand", "rosy", "tree"}

	fmt.Println()
	fmt.Println(boas)

	fmt.Printf("Length of boas is %d. Capacity of boas is %d.\n", len(boas), cap(boas))

	snakes := append(pythons, boas...)

	fmt.Println()
	fmt.Println(snakes)

	fmt.Printf("Length of snakes is %d. Capacity of snakes is %d.\n\n", len(snakes), cap(snakes))

	lesssnakes := append(snakes[1:3], snakes[5:8]...)
	fmt.Println(lesssnakes)
	fmt.Printf("Length of lesssnakes is %d. Capacity of lesssnakes is %d.\n\n", len(lesssnakes), cap(lesssnakes))

	fmt.Println("Now let's check this with integers")

	ones := []int{1, 2, 3, 4, 5, 6}

	tens := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	onestens := append(ones, tens...)

	fmt.Println()

	fmt.Printf("Length of ones is %d.\t Capacity of ones is %d.\n\n", len(ones), cap(ones))

	fmt.Printf("Length of tens is %d.\t Capacity of tens is %d.\n\n", len(tens), cap(tens))

	fmt.Printf("Length of onestens is %d.\t Capacity of onestens is %d.\n", len(onestens), cap(onestens))
}
