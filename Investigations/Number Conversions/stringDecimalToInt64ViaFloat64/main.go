package main

import (
	"fmt"
	"strconv"
)

//DecimalToInt64 - converts a decimal amount of money to an int64
func DecimalToInt64(decimal string) int64 {

	myFloat64, _ := strconv.ParseFloat(decimal, 64)
	fmt.Printf("myFloat64 is %v\n\n", myFloat64)
	myFloat64 *= 100
	myInt64 := int64(myFloat64)
	return myInt64
}

func main() {

	fmt.Println(DecimalToInt64("0.125"))

}
