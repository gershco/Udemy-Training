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

type WomanList struct {
	Data []Woman `json:"data"`
}

type Woman struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
	Balance   float64
}

var women WomanList

func main() {

	http.HandleFunc("/table", table)

	//http.HandleFunc("/withzeros", withzeros)

	http.HandleFunc("/male", male)

	http.HandleFunc("/female", female)

	http.ListenAndServe(":7070", nil)

}
func table(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "users.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func male(w http.ResponseWriter, r *http.Request) {

	m1 := Man{1, 25, "John", "Minor", "john@minor.gov", 125.38}
	m2 := Man{2, 34, "James", "Riley", "james@riley.com", 78.56}
	m3 := Man{3, 46, "Nigel", "Kennedy", "nigel@kennedy.mus", 21.59}

	men := ManList{[]Man{m1, m2, m3}}

	json, _ := json.Marshal(men)

	w.Write(json)

	//Xuser = nil

}

func female(w http.ResponseWriter, r *http.Request) {

	w1 := Woman{1, 25, "Jane", "Major", "jane@major.gov", 143.67}
	w2 := Woman{2, 34, "Julia", "Rigby", "julia@rigby.com", 35.99}
	w3 := Woman{3, 46, "Nicola", "Kenwood", "nicola@kenwood.mus", 21.59}

	women := WomanList{[]Woman{w1, w2, w3}}

	json, _ := json.Marshal(women)

	w.Write(json)

	//Xuser = nil

}
