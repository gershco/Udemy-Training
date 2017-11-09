package edit

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

func Edit(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		t, err := template.ParseFiles("edit.html", "note.html")
		if err != nil {
			panic(err)
		}

		err = t.ExecuteTemplate(w, "edit.html", nil)
		if err != nil {
			panic(err)
		}

	} else {

		Entry := r.FormValue("entry")

		//fmt.Println(Entry)

		note := []byte(Entry)

		ioutil.WriteFile("note.html", note, 0664)

	}
}
