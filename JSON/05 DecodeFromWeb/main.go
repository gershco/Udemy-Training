// Encoding means sending to a writer. NewEncoder wants a writer.
// Therefore we are going to send this to a web page.

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type person struct {
	First string
	Last  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/send", send)
	http.HandleFunc("/catch", catch)
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {

	p1 := person{"James", "Bond"}

	p2 := person{"Miss", "Moneypenny"}

	// Create slice of person
	xp := []person{p1, p2}

	// NewEncoder returns a new encoder that writes to w.
	// func NewEncoder(w io.Writer) *Encoder {

	// func (enc *Encoder) Encode(v interface{}) error {
	// Encode returns an error
	err := json.NewEncoder(w).Encode(xp)
	if err != nil {
		fmt.Println(err)
	}
}

func send(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "send.gohtml", nil)
}

var xp []person

func catch(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&xp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(xp)

	tpl.ExecuteTemplate(w, "catch.gohtml", xp)

}
