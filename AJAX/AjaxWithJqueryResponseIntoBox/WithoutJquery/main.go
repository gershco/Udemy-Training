// Minimal Client/Server AJAX Communication using golang web-server and JQuery
// Visit: http://127.0.0.1:8080
// Jquery script written by MAS
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// Index Request Handler
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "myajax.html", nil)
}

// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I was called by AJAX.")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ajax", ajaxHandler)
	http.ListenAndServe(":8080", nil)
}
