package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
}

type square struct {
	side float64
}

func (s square) Area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(sh shape) {
	fmt.Println(sh.Area())
}

func main() {

	sq1 := square{7}

	cir1 := circle{4}

	info(sq1)

	info(cir1)
}
