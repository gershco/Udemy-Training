package main

import (
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
	http.ListenAndServe(":7070", nil)

}

func grab(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, `
		<h1> Note Collection Function</h1>	
		<form method="POST">
		<textarea name="input"rows="15"cols="50"></textarea>
		<input type="submit" value="Submit note">
		</form>
		<br>`)

	} else {

		Entry = r.FormValue("input")

		Note := []byte(Entry)

		ioutil.WriteFile("note", Note, 0664)

		io.WriteString(w, "<b>Thank you. Your note has been received.</b>")

	}
}

func edit(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, ""+Entry)

}

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, "<b>You need to browse to localhost:7070/grab or localhost:7070/edit<b>")

}
