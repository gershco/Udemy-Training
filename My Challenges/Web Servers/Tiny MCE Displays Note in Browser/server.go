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

		Entry := r.FormValue("entry")

		fmt.Println(Entry)

		note := []byte(Entry)

		ioutil.WriteFile("note.html", note, 0664)

	}
}

func edit(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//io.WriteString(w, "Please check your html note."+Entry)

	http.ServeFile(w, r, "note.html")

	fmt.Println(Entry)
}

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, "<b>You need to browse to localhost:8081/grab or localhost:8081/edit<b>")
}
