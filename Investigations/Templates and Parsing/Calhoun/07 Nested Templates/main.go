package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

type Widget struct {
	Name  string
	Price int
}

type ViewData struct {
	Name    string
	Widgets []Widget
}

func main() {
	var err error
	testTemplate, err = template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := testTemplate.Execute(w, ViewData{
		Name: "John Smith",
		Widgets: []Widget{
			Widget{"Blue Widget", 15},
			Widget{"Red Widget", 13},
			Widget{"Green Widget", 12},
			Widget{"Yellow Widget", 14},
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
