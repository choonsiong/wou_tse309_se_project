package models

import (
	"context"
	"time"
)

type Publisher struct {
	ID            int       `json:"id"`
	PublisherName string    `json:"publisher_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// All returns a list of all publishers
func (p *Publisher) All() ([]*Publisher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, publisher_name, created_at, updated_at FROM publishers ORDER BY id`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var publishers []*Publisher

	for rows.Next() {
		var publisher Publisher
		err := rows.Scan(&publisher.ID, &publisher.PublisherName, &publisher.CreatedAt, &publisher.UpdatedAt)
		if err != nil {
			return nil, err
		}
		publishers = append(publishers, &publisher)
	}

	return publishers, nil
}

// GetByID get one publisher by id
func (p *Publisher) GetByID(id int) (*Publisher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, publisher_name, created_at, updated_at FROM publishers WHERE id = $1`
	row := db.QueryRowContext(ctx, query, id)

	var publisher Publisher

	err := row.Scan(
		&publisher.ID,
		&publisher.PublisherName,
		&publisher.CreatedAt,
		&publisher.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &publisher, nil
}

// Insert inserts a new publisher
func (p *Publisher) Insert(publisher Publisher) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `INSERT INTO publishers (publisher_name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id`

	var newID int

	err := db.QueryRowContext(ctx, stmt, publisher.PublisherName, time.Now(), time.Now()).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}
