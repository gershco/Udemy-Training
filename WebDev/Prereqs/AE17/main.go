package main

import "fmt"

func main() {

	s := "i'm sorry dave i can't do that"

	fmt.Println(s)

	fmt.Print()

	slices := []byte(s)

	fmt.Println(slices)

	stringslices := string(slices)

	fmt.Println(stringslices)

	fmt.Println(string(slices[:14]))

	fmt.Println(string(slices[17:]))

	for i := 0; i < len(slices); i++ {

		fmt.Println(string(slices[i]))
	}

	for _, v := range slices {

		fmt.Println(string(v))
	}
}
