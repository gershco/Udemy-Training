package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("td.html", td)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "table.gohtml", nil)
}

func td(w http.ResponseWriter, r *http.Request) {

	t := time.Now()
	h := t.Hour()
	min := t.Minute()
	d := t.Day()
	mon := t.Month()
	y := t.Year()

	date := fmt.Sprintf("Today's date is %d %s %d.", d, mon, y)
	time := fmt.Sprintf("The time now is %d:%d.", h, min)

	type Dt struct {
		Date string
		Time string
	}

	Tdnow := Dt{date, time}

	fmt.Println(Tdnow)

	err2 := tpl.ExecuteTemplate(w, "table.gohtml", Tdnow)
	if err2 != nil {
		log.Fatalln("template didn't execute: ", err2)
	}

}
