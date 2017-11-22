package main

import "fmt"

type gator int

type flamingo bool

type otter string

type swampCreature interface {
	greeting()
}

func (g gator) greeting() {

	fmt.Println("Hello, I am a gator.")
}

func (f flamingo) greeting() {

	fmt.Println("Hello, I am pink, beautiful and wonderful.")
}

func (o otter) greeting() {

	fmt.Println("Hello, I'm sleek and swift and love catching fish!")
}

func bayou(s swampCreature) {

	s.greeting()
}

func main() {

	var gatee gator
	bayou(gatee)

	var flamee flamingo
	bayou(flamee)

	var ottee otter
	bayou(ottee)
}
