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

// GetByID gets one author by id
func (a *Author) GetByID(id int) (*Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, authors.author_name, created_at, updated_at FROM authors WHERE id = $1`
	row := db.QueryRowContext(ctx, query, id)

	var author Author

	err := row.Scan(
		&author.ID,
		&author.AuthorName,
		&author.CreatedAt,
		&author.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &author, nil
}

// Insert inserts a new author
func (a *Author) Insert(author Author) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `INSERT INTO authors (author_name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id`

	var newID int

	err := db.QueryRowContext(ctx, stmt, author.AuthorName, time.Now(), time.Now()).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

// DeleteForBookId delete records for given book id
func (a *Author) DeleteForBookId(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM mybooks_authors WHERE book_id = $1`
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
