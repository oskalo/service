package main

import (
	"database/sql"
	"log"
)

func main() {
	//worst practice ever
	connStr := "sslmode=verify-full port=5432 user=user dbname=example password=test"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err  = db.Query("migration query")
	if err != nil {
		log.Fatal(err)
	}
}
