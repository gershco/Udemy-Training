package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("letter.phpasp")

	if err != nil {
		fmt.Println("There was an error parsing", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("There was an error executing the template", err)
	}
}
