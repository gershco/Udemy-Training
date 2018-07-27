package main

import (
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

}
