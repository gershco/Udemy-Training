package main

import (
	"fmt"
)

func main() {

	// Remember: A string is a collection of runes and runes are bytes. Therefore strings are a slice of bytes.

	first := rune("Go"[0]) //Convert position zero of the string into a rune.

	second := rune("Go"[1]) //Convert position one of the string into a rune.

	fmt.Println(first, second)
	fmt.Println()

	word := "cantankerous"

	for i := 0; i < len(word); i++ {

		fmt.Printf("%v \t %T \n", rune(word[i]), rune(word[i]))
	}
}
