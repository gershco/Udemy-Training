package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

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

func main() {

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

	type User struct {
		ID        int    `json:"userid"`
		Age       int    `json:"userage"`
		FirstName string `json:"userfirst"`
		LastName  string `json:"userlast"`
		Email     string `json:"useremail"`
	}

	sqlStatement := `SELECT * FROM users WHERE id=$1;`
	var user User
	row := db.QueryRow(sqlStatement, 11)

	switch err := row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email); err {

	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println("user struct is: ", user)
	default:
		panic(err)
	}

	userjson, err := json.Marshal(user) // json.Marshal returns the JSON encoding of emp as a []byte.
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(userjson))
}
