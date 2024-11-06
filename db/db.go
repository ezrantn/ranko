package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const file string = "search.db"

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	return db
}
