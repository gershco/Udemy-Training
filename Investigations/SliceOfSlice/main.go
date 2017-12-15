package main

import "fmt"

func main() {

	f := "Freilichen"
	g := "Chanukah"

	sf := []byte(f)
	sg := []byte(g)

	slices := [][]byte{sf, sg}

	fmt.Println(slices)

}
