package main

import (
	"fmt"
	"net/http"

	"Github.com/gershco/UdemyTraining/Notes/edit"
	"Github.com/gershco/UdemyTraining/Notes/grab"
)

//Entry is used by different functions.
var Entry string

func main() {
	http.HandleFunc("/grab", grab.Grab)
	http.HandleFunc("/edit", edit.Edit)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", nil)

	fmt.Println("Serving on localhost:8081/grab and /edit")
}
