package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheel bool
	role      string
}

type sedan struct {
	vehicle
	luxury bool
	role   string
}

func (tr truck) transportDevice() {
	fmt.Printf("All day, I %s.\n", tr.role)
}

func (se sedan) transportDevice() {
	fmt.Printf("All day, I %s.\n", se.role)
}

func main() {

	t := truck{vehicle{3, "red"}, true, "transport soil around the country"}

	s := sedan{vehicle{5, "black"}, false, "tootle around town"}

	t.transportDevice()
	s.transportDevice()

}
