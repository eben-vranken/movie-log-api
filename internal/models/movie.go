package models

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	GenreID     int       `json:"genre_id"`
	Director    string    `json:"director"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	Rating      float32   `json:"rating"`
}
