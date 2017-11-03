package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/grab", grab)
	http.HandleFunc("/edit", edit)

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

		entry := r.FormValue("input")

		fmt.Println(entry)

		note := []byte(entry)

		ioutil.WriteFile("note", note, 0664)

		io.WriteString(w, "<b>Thank you. Your note has been received.</b>")

	}
}

func edit(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "note")

}
