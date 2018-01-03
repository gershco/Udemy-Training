// Encoding means sending to a writer. NewEncoder wants a writer.
// Therefore we are going to send this to a web page.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	First string
	Last  string
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {

	p1 := person{"James", "Bond"}

	p2 := person{"Miss", "Moneypenny"}

	// Create slice of person
	xp := []person{p1, p2}

	fmt.Println(xp)

	// NewEncoder returns a new encoder that writes to w.
	// func NewEncoder(w io.Writer) *Encoder {

	// func (enc *Encoder) Encode(v interface{}) error {
	// Encode returns an error
	err := json.NewEncoder(w).Encode(xp)
	if err != nil {
		fmt.Println(err)
	}
}
