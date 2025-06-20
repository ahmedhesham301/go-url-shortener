package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "urls.db")
	if err != nil {
		panic("couldn't connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createUrlsTable := `
	CREATE TABLE IF NOT EXISTS urls (
	  urlHash INTEGER PRIMARY KEY AUTOINCREMENT,
	  url TEXT NOT NULL
	)
	`

	_,err := DB.Exec(createUrlsTable)
	if err != nil {
		panic("could not create table "+ err.Error())
	}
}