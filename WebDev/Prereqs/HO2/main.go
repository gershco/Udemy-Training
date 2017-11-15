package main

import (
	"fmt"
)

type person struct {
	firstName, lastName string
}
type secret struct {
	who          person
	code, weapon string
}

func (p person) pSpeak() {
	fmt.Println(p.firstName, p.lastName, `says, "Good morning, James."`)
}

func (s secret) sSpeak() {
	fmt.Println(s.who.firstName, `replies, "You're out of luck, Moneypenny. My`, s.weapon, "is loaded and", s.code, `is ready for action!"`)
}
func main() {

	p1 := person{"Miss", "Moneypenny"}

	p1.pSpeak()

	s1 := secret{person{"James", "Bond"}, "007", "Walther PPK"}

	s1.sSpeak()
}
