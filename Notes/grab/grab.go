package grab

import (
	"io/ioutil"
	"net/http"
)

func Grab(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		http.ServeFile(w, r, "grab.html")

	} else {

		Entry := r.FormValue("entry")

		//fmt.Println(Entry)

		note := []byte(Entry)

		ioutil.WriteFile("note.html", note, 0664)

	}
}
