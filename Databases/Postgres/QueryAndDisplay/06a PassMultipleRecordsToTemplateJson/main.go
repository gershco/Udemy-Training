// Taken from: https://www.calhoun.io/querying-for-multiple-records-with-gos-sql-package/

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

// Declare constants for db connection.

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1025"
	dbname   = "postgres"
)

var templates = template.Must(template.ParseGlob("templates/*"))

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
	Balance   float64
}

type UserList struct {
	Data []User `json:"data"`
}

var person User

var Xuser []User

func main() {

	http.HandleFunc("/users", users)

	http.HandleFunc("/ajaxcall", callJSON)

	http.ListenAndServe(":8081", nil)

}
func users(w http.ResponseWriter, r *http.Request) {

	// Create string containing all db connection information.

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Validate the connection arguments provided.

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Open a connection to the db in order to validate connection success.

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {

		/*
			As Query() returns an error in addition to the *Rows, we need to handle this error
			before doing anything else.
		*/

		panic(err)
	}

	defer rows.Close()

	/*
		rows.Next starts a for loop which iterates over each record returned by the
		SQL statement. It returns true when the next row is successfully prepared,
		and false otherwise.
	*/

	for rows.Next() {

		err = rows.Scan(&person.ID, &person.Age, &person.FirstName, &person.LastName, &person.Email, &person.Balance)
		if err != nil {
			// handle this error
			panic(err)
		}

		Xuser = append(Xuser, person)
	}

	err1 := templates.ExecuteTemplate(w, "users.html", nil)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func callJSON(w http.ResponseWriter, r *http.Request) {

	data := UserList{Xuser}

	json, err := json.Marshal(data)

	if err != nil {
		fmt.Println("The struct could not be marshalled into JSON.")
	}

	w.Write(json)

	Xuser = nil

}
