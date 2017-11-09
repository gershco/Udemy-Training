package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	b, _ := ioutil.ReadFile("note.html")

	s := string(b)

	fmt.Print(b)

	fmt.Print()

	fmt.Println(s)
}
