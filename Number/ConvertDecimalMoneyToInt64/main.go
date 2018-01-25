package main

import (
	"fmt"
	"strconv"
)

//DecimalMoneyToInt64 converts a decimal amount of money to an int64
func DecimalMoneyToInt64(decimalMoney string) int64 {

	myFloat64, _ := strconv.ParseFloat(decimalMoney, 64)
	myFloat64 *= 100
	myInt64 := int64(myFloat64)
	return myInt64
}

func main() {

	converted := DecimalMoneyToInt64("14.34")

	fmt.Printf("converted, of type %T, is %v", converted, converted)
}
