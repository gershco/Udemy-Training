package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type Crimes struct {
	Datetime    string
	Address     string
	District    string
	Beat        string
	Grid        string
	Description string
	Code        string
	Latitude    string
	Longitude   string
}

func main() {

	http.HandleFunc("/", crime)

	http.ListenAndServe(":8081", nil)

}

func crime(w http.ResponseWriter, r *http.Request) {

	csvFile, err1 := os.Open("crime.csv")
	if err1 != nil {
		fmt.Println("The file could not be opened.")
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var details []Crimes
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		details = append(details, Crimes{
			Datetime:    line[0],
			Address:     line[1],
			District:    line[2],
			Beat:        line[3],
			Grid:        line[4],
			Description: line[5],
			Code:        line[6],
			Latitude:    line[7],
			Longitude:   line[8],
		},
		)
	}

	t, _ := template.ParseFiles("crime.html")

	err := t.Execute(w, details)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}

}
