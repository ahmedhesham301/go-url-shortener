package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var createUrlsTable = `
	CREATE TABLE IF NOT EXISTS urls (
	  id SERIAL PRIMARY KEY,
	  url VARCHAR NOT NULL
	)
	`

func InitDB() {
	pool, err := pgxpool.New(context.Background(), "postgres://ahmed:1234@localhost:5432/postgres")
	if err != nil {
		panic("couldn't connect to database")
	}
	defer conn.Close(context.Background())
	// DB.SetMaxOpenConns(10)
	// DB.SetMaxIdleConns(5)

	// Create table

	_, err = conn.Exec(context.Background(), createUrlsTable)
	if err != nil {
		panic("could not create table " + err.Error())
	}
}
