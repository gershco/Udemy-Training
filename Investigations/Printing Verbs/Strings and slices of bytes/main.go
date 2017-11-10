package main

import (
	"fmt"
)

func main() {

	s := "Where did you get that hat?"

	fmt.Println("String s is: ", s)

	fmt.Print("\n%s gives:")

	fmt.Printf("\t%s", s)

	fmt.Print("\n\n%q gives:")

	fmt.Printf("\t%q", s)

	fmt.Print("\n\n%x gives:")

	fmt.Printf("\t%x", s)

	fmt.Print("\n\n%X gives:")

	fmt.Printf("\t%X", s)

}
