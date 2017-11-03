package main

import (
	"fmt"
)

//interface definition
type vowelsFinder interface {
	findVowels() []rune
}

type myString string

//myString implements vowelsFinder
func (ms myString) findVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := myString("Sam Anderson")
	var v vowelsFinder
	v = name // possible since myString implements vowelsFinder
	fmt.Printf("Vowels are %c", v.findVowels())

}
