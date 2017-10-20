package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := 128

	b := "\n\n2^7"

	fmt.Print("\na, of type ")

	fmt.Printf("%T,", a) // Print to standard output the type of a

	fmt.Print(" is assigned the value ")

	fmt.Print(a, ".")

	fmt.Print("\n\nstrconv.Itoa(a) is of type ")

	fmt.Printf("%T", strconv.Itoa(a)) // Print to standard output the type of strconv.Itoa(a)

	fmt.Println(b + " = " + strconv.Itoa(a)) // Concatenate b and the converted integer

	c := "25"

	d := "37"

	e, _ := strconv.Atoi(c) // strconv.Atoi has two returns, so a blank identifier is required

	f, _ := strconv.Atoi(d)

	fmt.Print("\n\n25 + 37 returns ")

	fmt.Print(c + d)

	fmt.Print(" as c and d are strings")

	fmt.Print("\n\n25 + 37 returns ")

	fmt.Print(e + f)

	fmt.Print(" as e and f are integers returned from strconv.Atoi")

}
