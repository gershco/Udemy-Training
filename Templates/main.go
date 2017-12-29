package main

import (
	"html/template"
	"net/http"
)

// Student is exported struct with exported fields
type Student struct {
	Name string
	Age  int
}

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))

func main() {

	http.HandleFunc("/", index)

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	http.ListenAndServe(":7070", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	s := Student{"Joe Bloggs", 34}

	// you access the cached templates with the defined name, not the filename
	err := templates.ExecuteTemplate(w, "indexPage", s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
