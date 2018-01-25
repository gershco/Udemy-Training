package main

import (
	"fmt"
	"time"
)

func main() {

	layout := "02111972"
	str := "02012006"
	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nt, which is of type %T is %v\n\n", t, t)

	value := "Thu, 05/19/11, 10:47PM"
	// Writing down the way the standard time would look like formatted our way
	layout1 := "Mon, 01/02/06, 03:04PM"
	t1, _ := time.Parse(layout1, value)
	fmt.Printf("t1, which is of type %T is %v\n\n", t1, t1)

	dtstr1 := "1972-11-02"
	dt, _ := time.Parse("2006-01-02", dtstr1)
	dtstr2 := dt.Format("02012006")

	fmt.Printf("dt, which is of type %T is %v\n\n", dt, dt)

	fmt.Printf("dtstr2 is of type %T is %v\n\n", dtstr2, dtstr2)

}
