package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func Connect() *sql.DB {
	var err error
	db, err := sql.Open("sqlite", "file:database.db?_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successful connection to DataBase")
	}

	return db
}

func CloseDB() {
	if Connect() != nil {
		Connect().Close()
	}
}
