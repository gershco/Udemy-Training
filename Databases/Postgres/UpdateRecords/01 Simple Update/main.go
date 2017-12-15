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

	sqlUpdate := `  
	UPDATE users  
	SET first_name = $2, last_name = $3, email = $4   
	WHERE id = $1;`
	_, err = db.Exec(sqlUpdate, 4, "Pink", "Sky", "pink@sky.io")
	if err != nil {
		panic(err)
	}

	sqlAgeUpdate := `  
	UPDATE users  
	SET age = $2
	WHERE id = $1;`
	_, err = db.Exec(sqlAgeUpdate, 1, 48)
	if err != nil {
		panic(err)
	}

}
