package main

import (
	"database/sql"
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

	/*
	   Could insert a record like this:

	   sqlStatement := `
	   INSERT INTO users (age, email, first_name, last_name)
	   VALUES ($1, $2, $3, $4)`
	   _, err = db.Exec(sqlStatement, 45, "fred@bloggs.com", "Fred", "Bloggs")
	   if err != nil {
	     panic(err)
	   }

	   but we need to be able to get the ID of the newly-created record.

	   Here is a better method:
	*/
	// Insert records. Use 'RETURNING ID' to return the id of the newly-inserted record.

	// Use db.QueryRow() as we only expect to get one row back.

	// Scan the returned Row and store the ID in the iteger variable id.

	sqlStatement := `  
	INSERT INTO users (age, email, first_name, last_name)  
	VALUES ($1, $2, $3, $4)  
	RETURNING *;`
	var id, age int
	var email, firstName, lastName string
	err = db.QueryRow(sqlStatement, 26, "clarisa@mariposa.es", "Clarisa", "Mariposa").Scan(&id, &age, &firstName, &lastName, &email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID \t %d \n Age \t %d \n First Name \t %s \n Last Name \t %s \n Email \t %s \n", id, age, firstName, lastName, email)

}
