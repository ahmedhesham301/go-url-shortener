package models

import (
	"context"
	"urlshortener/db"
)

type Url struct {
	Id  int    `json:"id"`
	URL string `json:"url"`
}

func (u *Url) Save() error {
	query := `
	INSERT INTO urls(url)
	VALUES ($1)
	RETURNING id
	`
	err := db.Pool.QueryRow(context.Background(), query, u.URL).Scan(&u.Id)
	return err
}

func GetUrlById(id string) (*Url, error) {
	query := "SELECT * FROM urls WHERE id = ($1)"
	row := db.Pool.QueryRow(context.Background(), query, id)

	var url Url
	err := row.Scan(&url.Id, &url.URL)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
