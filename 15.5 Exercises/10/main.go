package main

import "fmt"

func main() {

	b := "Where did you get that hat?"

	fmt.Print("String b is: \t \t \t")
	fmt.Println(b)

	fmt.Print("It is obviously a \t \t")
	fmt.Printf("%T \n", b)

	fmt.Print("The rune at position 2 is: \t")
	fmt.Println(b[2])

	fmt.Print("It is of type: \t \t \t")
	fmt.Printf("%T \n", b[2])

}
