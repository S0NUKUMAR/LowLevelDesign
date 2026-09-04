package domain

import (
	"time"
)

type Movie struct {
	ID string
	Title string
	Language string
	Duration time.Duration
	Genre string
}


func NewMovie(id, title, language, genre string , duration time.Duration) *Movie {
	return &Movie{
		ID:id,
		Title:title,
		Language:language,
		Genre:genre,
		Duration:duration,
	}
}


