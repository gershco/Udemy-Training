// Taken from: https://www.calhoun.io/querying-for-multiple-records-with-gos-sql-package/

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

	rows, err := db.Query("SELECT * FROM users order by id;")
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

	type User struct {
		ID        int
		Age       int
		FirstName string
		LastName  string
		Email     string
	}

	var user User

	xuser := []User{}

	for rows.Next() {

		err = rows.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(user)

		xuser = append(xuser, user)

	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println()

	fmt.Println(xuser)

}
