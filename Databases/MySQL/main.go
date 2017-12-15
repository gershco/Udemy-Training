package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var (
		accountID    int
		availBalance float64
	)
	rows, err := db.Query("select account_id, avail_balance from account where account_id = ?", 3)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&accountID, &availBalance)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(accountID, availBalance)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
