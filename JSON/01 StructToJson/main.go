package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name   string `json:"empname"`
	Number int    `json:"empid"`
}

func main() {
	emp := Employee{Name: "Rocky", Number: 5454}
	e, err := json.Marshal(emp) // json.Marshal returns the JSON encoding of emp as a []byte.
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("emp is: ", emp) //emp is the original struct
	fmt.Println()
	fmt.Println("e is: ", e) //e is the []byte of the JSON encoded emp.
	fmt.Println()
	fmt.Println("string(e) is :", string(e))
}
