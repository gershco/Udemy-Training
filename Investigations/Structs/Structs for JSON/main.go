package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*"))

type ManList struct {
	Data []Man `json:"data"`
}

type Man struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
	Balance   float64
}

var men ManList

func main() {

	http.HandleFunc("/table", table)

	http.HandleFunc("/ajaxcall", callJSON)

	http.ListenAndServe(":7070", nil)

}
func table(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "users.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func callJSON(w http.ResponseWriter, r *http.Request) {

	m1 := Man{1, 25, "John", "Minor", "john@minor.gov", 125.38}
	m2 := Man{2, 34, "James", "Riley", "james@riley.com", 78.56}
	m3 := Man{3, 46, "Nigel", "Kennedy", "nigel@kennedy.mus", 21.59}

	men := ManList{[]Man{m1, m2, m3}}

	json, _ := json.Marshal(men)

	w.Write(json)
}
