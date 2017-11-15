package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{vehicle{3, "red"}, true}

	s := sedan{vehicle{5, "black"}, false}

	fmt.Println(t)

	fmt.Println(s)

	fmt.Println(t.colour)

	fmt.Println(s.doors)

}
