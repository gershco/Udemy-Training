package main

import (
	"fmt"
)

func main() {

	vegarray := [...]string{"tomato", "cucumber", "pepper", "courgette", "cauliflower", "brocolli"}

	fmt.Println("vegarray is", vegarray)
	fmt.Println()

	vegslice := vegarray[1:3]
	fmt.Println(`vegslice, formed by "vegarray[1:3]" is`, vegslice)

	fmt.Printf("Length of vegslice is %d. Capacity is %d\n", len(vegslice), cap(vegslice))
	fmt.Println()

	vegslice = vegslice[:cap(vegslice)] //re-slicing vegslice till its capacity
	fmt.Println(`Reslicing with "vegslice = vegslice[:cap(vegslice)]" produces vegslice:`, vegslice)
	fmt.Println("After re-slicing length is", len(vegslice), "and capacity is", cap(vegslice))

}
