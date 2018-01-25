package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	d := t.Day()
	mon := t.Month()
	y := t.Year()

	h := t.Hour()
	min := t.Minute()

	time := fmt.Sprintf("The time now is %d:%d.", h, min)

	date := fmt.Sprintf("Today's date is %d %s %d.", d, mon, y)

	fmt.Println(time)
	fmt.Println(date)

}
