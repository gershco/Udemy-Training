package main

import (
	"fmt"
)

type person interface {
	speak()
}

type civvy struct {
	name string
}
type secret struct {
	name string
}

func (c civvy) speak() {
	fmt.Println(c.name, "would love to be noticed.")
}

func (s secret) speak() {
	fmt.Printf("Unfortunately, %s's eyes are always looking elsewhere!\n\n", s.name)
}

func talk(p person) {
	p.speak()
}

func main() {

	p1 := civvy{"Miss Moneypenny"}
	s1 := secret{"James"}

	talk(p1)
	talk(s1)
}
