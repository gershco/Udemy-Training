package main

import "fmt"

type gator int

func (g gator) greeting() {

	fmt.Println("Hello, I am a gator.")
}

func main() {

	var gatee gator

	gatee.greeting()
}
