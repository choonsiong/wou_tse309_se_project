package models

import "time"

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	PublisherID     int       `json:"publisher_id"`
	PublicationYear int       `json:"publication_year"`
	Description     string    `json:"description"`
	Slug            string    `json:"slug"`
	Authors         []Author  `json:"authors"`
	AuthorIDs       []int     `json:"author_ids,omitempty"`
	Genres          []Genre   `json:"genres"`
	GenreIDs        []int     `json:"genre_ids,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
