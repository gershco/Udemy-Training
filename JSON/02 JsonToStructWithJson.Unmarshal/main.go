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

	j := `[{"First":"James","Last":"Bond"},{"First":"Miss","Last":"Moneypenny"}]`
	fmt.Println("json:", j)
	fmt.Println()

	// Create slice of person because that is what j actually is.
	xp := []person{}

	// func Unmarshal(data []byte, v interface{}) error {
	// Unmarshal takes a slice of bytes and stores the result
	// in the value pointed to by v. It returns an error.
	err := json.Unmarshal([]byte(j), &xp)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Go data: %+v\n\n", xp)

	for i, v := range xp {

		fmt.Println(i, v, "\t", v.First, "\t", v.Last)

	}

}
