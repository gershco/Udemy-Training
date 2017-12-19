package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/time", timee)
	http.HandleFunc("/date", datee)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "table.gohtml", nil)
}

func timee(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	h := t.Hour()
	min := t.Minute()
	time := fmt.Sprintf("The time now is %d:%d.", h, min)
	fmt.Fprintln(w, time)
}

func datee(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	d := t.Day()
	mon := t.Month()
	y := t.Year()
	date := fmt.Sprintf("Today's date is %d %s %d.", d, mon, y)
	fmt.Fprintln(w, date)
}
