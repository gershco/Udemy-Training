package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	role string
}

type sedan struct {
	vehicle
	role string
}

func (a truck) transportationDevice() string {
	return fmt.Sprintln("All day, I", a.role)
}

func (b sedan) transportationDevice() string {
	return fmt.Sprintln("I am a", b.doors, "door sedan. I spend my day", b.role)
}

type transportation interface {
	transportationDevice() string
}

func report(c transportation) {
	fmt.Println(c.transportationDevice())
}

func main() {

	d := truck{vehicle{3, "red"}, "transport soil around the country"}

	e := sedan{vehicle{5, "black"}, "tootling around town"}

	report(d)

	report(e)
}
