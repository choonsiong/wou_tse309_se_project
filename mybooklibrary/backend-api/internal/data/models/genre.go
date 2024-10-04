package models

import (
	"context"
	"time"
)

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// All returns a list of all genres
func (g *Genre) All() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, genre_name, created_at, updated_at FROM genres ORDER BY genre_name`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var genres []*Genre

	for rows.Next() {
		var genre Genre
		err := rows.Scan(&genre.ID, &genre.GenreName, &genre.CreatedAt, &genre.UpdatedAt)
		if err != nil {
			return nil, err
		}
		genres = append(genres, &genre)
	}

	return genres, nil
}

// GetByID get one genre by id
func (g *Genre) GetByID(id int) (*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, genre_name, created_at, updated_at FROM genres WHERE id = $1`
	row := db.QueryRowContext(ctx, query, id)

	var genre Genre

	err := row.Scan(
		&genre.ID,
		&genre.GenreName,
		&genre.CreatedAt,
		&genre.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &genre, nil
}

// Insert inserts a new genre
func (g *Genre) Insert(genre Genre) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `INSERT INTO genres (genre_name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id`

	var newID int

	err := db.QueryRowContext(ctx, stmt, genre.GenreName, time.Now(), time.Now()).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

// DeleteForBookId delete records for given book id
func (g *Genre) DeleteForBookId(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM mybooks_genres WHERE book_id = $1`
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
