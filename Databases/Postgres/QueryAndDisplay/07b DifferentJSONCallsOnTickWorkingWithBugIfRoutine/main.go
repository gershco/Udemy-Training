package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*"))

type Person struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
	Balance   float64
}

type List struct {
	Data []Person `json:"data"`
}

var person Person

var People []List

func main() {

	http.HandleFunc("/table", table)

	http.HandleFunc("/data", data)

	http.ListenAndServe(":7070", nil)

}
func table(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "users.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func data(w http.ResponseWriter, r *http.Request) {

	var check bool

	//check = true

	fmt.Printf("%t", check)

	if check {

		p1 := Person{1, 25, "Jane", "Major", "jane@major.gov", 143.67}
		p2 := Person{2, 34, "Julia", "Rigby", "julia@rigby.com", 35.99}
		p3 := Person{3, 46, "Nicola", "Kenwood", "nicola@kenwood.mus", 21.59}

		people := List{[]Person{p1, p2, p3}}

		json, _ := json.Marshal(people)

		w.Write(json)

		fmt.Println(string(json))

		People = nil

	} else {

		p1 := Person{1, 25, "John", "Minor", "john@minor.gov", 125.38}
		p2 := Person{2, 34, "James", "Riley", "james@riley.com", 78.56}
		p3 := Person{3, 46, "Nigel", "Kennedy", "nigel@kennedy.mus", 21.59}

		people := List{[]Person{p1, p2, p3}}

		json, _ := json.Marshal(people)

		w.Write(json)

		fmt.Println(string(json))

		People = nil

	}

}
