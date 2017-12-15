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

	//First, construct the SQL statement.
	sqlStatement := `SELECT id, email FROM users WHERE id=$1;`

	var email string
	var id int

	/* Use the QueryRow() method. Pass the SQL statement as the first argument,
	followed by data for the construction of the SQL statment as additional arguments. */
	row := db.QueryRow(sqlStatement, 12)
	// row is a *Row ( a pointer to a sql.Row)

	/* Next line calls row.Scan() method. Tells Scan() to copy the retrieved data into the
	memory locations used by id and email. Scan() returns nil if this successful, otherwise
	it returns an error.*/

	switch err := row.Scan(&id, &email); err {
	// First case - No rows returned.
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	// Second case - No error.
	case nil:
		fmt.Println(id, email)
	// Third case - There is an error, but not the ErrNoRows error.
	default:
		panic(err)
	}

}
