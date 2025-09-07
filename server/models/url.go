package models

import (
	"urlshortener/db"
)

type Url struct {
	Id string `json:"id"`
	URL     string `json:"url"`
}

func (u Url) Save() error {
	query := `
	INSERT INTO urls(url)
	VALUES (?)
	`
	_, err := db.DB.Exec(query, u.URL)
	return err
}

func GetUrlById(id string) (*Url, error) {
	query := "SELECT * FROM urls WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var url Url
	err := row.Scan(&url.Id, &url.URL)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
