package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*"))

//************************************************************************************
//	WEBFORM TO INPUT INVOICE DATA
//************************************************************************************

func serveForm(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		err := templates.ExecuteTemplate(w, "test.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {

		//Parse the form in order to extract the invoice lines

		parseErr := r.ParseForm()

		if parseErr != nil {
			fmt.Println("Error parsing the form")
		}

		io.WriteString(w, "Data received successfully")

		//fmt.Printf("\n\n\n\n invoiceHeader, of type %T, is %+v\n\n", invoiceHeader, invoiceHeader)

	}
}

//************************************************************************************
//START PROGRAM
//************************************************************************************

func main() {

	fmt.Println("Serving on localhost:8080/sf")
	http.HandleFunc("/sf", serveForm)

	http.ListenAndServe(":8080", nil)

}
