package main

import "fmt"

func main() {

	fmt.Print("Please enter your name: ")

	var input string

	fmt.Scan(&input)

	fmt.Println ("Hello" , input)

}