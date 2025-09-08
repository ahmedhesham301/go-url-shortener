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
var Pool *pgxpool.Pool

func InitDB() {
	conn, err := pgxpool.ParseConfig("postgres://postgres:1234@localhost:5432/postgres")
	if err != nil {
		panic("could not parse conection string: " + err.Error())
	}
	conn.ConnConfig.TLSConfig = nil

	Pool, err = pgxpool.NewWithConfig(context.Background(), conn)
	if err != nil {
		panic("couldn't con	nect to database")
	}

	_, err = Pool.Exec(context.Background(), createUrlsTable)
	if err != nil {
		panic("could not create table " + err.Error())
	}
}
