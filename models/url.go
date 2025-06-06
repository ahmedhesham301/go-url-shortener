package models

import (
	"urlshortner/db"
)

type Url struct {
	UrlHash string `json:"urlHash"`
	URL     string `json:"url"`
}

func (u Url) Save() error {
	query := `
	INSERT INTO urls(urlHash, url)
	VALUES (?, ?)
	`
	_, err := db.DB.Exec(query, u.UrlHash, u.URL)
	return err
}

func GetUrlByHash(hash string) (*Url, error) {
	query := "SELECT * FROM urls WHERE urlHash = ?"
	row := db.DB.QueryRow(query, hash)

	var url Url
	err := row.Scan(&url.UrlHash, &url.URL)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
