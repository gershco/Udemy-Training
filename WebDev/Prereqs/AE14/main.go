package main

import "fmt"

type gator int

func main() {

	var g1 gator

	g1 = 4

	fmt.Println(g1)

	fmt.Printf("%T\n", g1)

	var ig1 int

	fmt.Println(ig1)

	fmt.Printf("%T\n", ig1)

	ig1 = int(g1)

	fmt.Println(ig1)

	fmt.Printf("%T\n", ig1)

}
