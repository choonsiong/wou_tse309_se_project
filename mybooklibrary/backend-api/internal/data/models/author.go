package models

import (
	"context"
	"time"
)

type Author struct {
	ID         int       `json:"id"`
	AuthorName string    `json:"author_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// All returns a list of all authors
func (a *Author) All() ([]*Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, author_name, created_at, updated_at  FROM authors ORDER BY author_name`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var authors []*Author

	for rows.Next() {
		var author Author
		err := rows.Scan(&author.ID, &author.AuthorName, &author.CreatedAt, &author.UpdatedAt)
		if err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}
	return authors, nil
}
