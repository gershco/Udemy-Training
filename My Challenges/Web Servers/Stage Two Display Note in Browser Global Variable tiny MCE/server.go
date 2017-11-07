package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

//Entry is used by different functions.
var Entry string

func main() {
	http.HandleFunc("/grab", grab)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/", index)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", nil)

}

func grab(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		http.ServeFile(w, r, "grab.html")

	} else {

		Entry := r.FormValue("input")

		fmt.Println(Entry)

		note := []byte(Entry)

		ioutil.WriteFile("note", note, 0664)

		//io.WriteString(w, "<b>Your note has been saved.</b>")
		//http.ServeFile(w, r, "saved.html")

	}
}

func edit(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, "<b>Please check your note. </b><br> "+Entry)

}

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, "<b>You need to browse to localhost:7070/grab or localhost:7070/edit<b>")

}
