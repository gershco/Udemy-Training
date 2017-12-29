package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
}

func main() {

	p1 := person{"James", "Bond"}

	p2 := person{"Miss", "Moneypenny"}

	// Create slice of person
	xp := []person{p1, p2}

	fmt.Println("Slice of person:", xp)

	fmt.Println()

	fmt.Printf("Go data: %+v\n\n", xp)

	// func Marshal(v interface{}) ([]byte, error) {
	// Marshal takes some data and returns a slice of byte and an error.
	bs, err := json.Marshal(xp)

	if err != nil {
		fmt.Println(err)
	}
	// Slice of bytes 'bs' needs to be converted to a string.
	fmt.Println("Json:", string(bs))
}
