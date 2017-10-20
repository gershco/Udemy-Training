package main

import "fmt"

func main() {

	a := []byte{87, 104, 97, 116, 32, 97, 114, 101, 32, 121, 111, 117, 32, 108, 111, 111, 107, 105, 110, 103, 32, 97, 116, 63}

	fmt.Println(string(a))

	fmt.Printf("\n%T", string(a))

}
