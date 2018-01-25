package main

import (
	"fmt"
	"strconv"
	"strings"
)

//DecimalToInt64 -
func DecimalToInt64(decimal string) int64 {

	var moneyInt64 int64

	dpPosition := strings.IndexByte(decimal, '.')

	difference := len(decimal) - dpPosition

	if dpPosition == -1 {

		moneyInt, _ := strconv.Atoi(decimal)

		moneyInt64 = int64(moneyInt) * 100

		fmt.Println(moneyInt64)

		return moneyInt64

	}

	switch difference {

	case 2:

		fmt.Println("Work in progress.")

		return 0
	}

	fmt.Println("The original amount contains a decimal point and needs further manipulation.")

	return 0

}

func main() {

	DecimalToInt64("21.9")
}

//fmt.Printf("Position of decimal point is %v\n\n", strings.IndexByte(stringDecimnalNumber, '.'))

//fmt.Printf("Length of stringDecimnalNumber is %v\n\n", len(stringDecimnalNumber))

/*

package main

import (
	"fmt"
	"math"
)

func main() {
var x,dp float64
x =1
dp = 3
y := x * math.Pow(10,dp)
	fmt.Println(y)
}


*/
