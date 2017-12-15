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

	sqlVerifyUpdateCount := `  
	UPDATE users  
	SET first_name = $2, last_name = $3, email = $4  
	WHERE id = $1;`
	res, err := db.Exec(sqlVerifyUpdateCount, 5, "Adam", "Aardvark", "adam@aardvark.com")
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

}
