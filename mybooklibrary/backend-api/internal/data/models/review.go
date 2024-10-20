package models

import (
	"context"
	"time"
)

type Review struct {
	ID        int       `json:"id"`
	Review    string    `json:"review"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// All returns all reviews
func (r *Review) All() ([]*Review, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, review, rating, created_at, updated_at FROM reviews ORDER BY id`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*Review

	for rows.Next() {
		var review Review
		err := rows.Scan(&review.ID, &review.Review, &review.Rating, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, nil
}

// Insert inserts a new record to both reviews and book_reviews tables
func (r *Review) Insert(review string, rating int, userID int, bookID int) (int, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	// Insert into review table first
	stmt := `INSERT INTO reviews (review, rating, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	var reviewID int

	err := db.QueryRowContext(ctx, stmt,
		review,
		rating,
		time.Now(),
		time.Now(),
	).Scan(&reviewID)
	if err != nil {
		return -1, -1, err
	}

	// Insert into book_reviews table then
	stmt2 := `INSERT INTO book_reviews (review_id, user_id, book_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var bookReviewID int

	err = db.QueryRowContext(ctx, stmt2,
		reviewID,
		userID,
		bookID,
		time.Now(),
		time.Now(),
	).Scan(&bookReviewID)
	if err != nil {
		return -1, -1, err
	}

	return reviewID, bookReviewID, nil
}

// Update updates book review
func (r *Review) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `UPDATE reviews SET review = $1, rating = $2, updated_at = $3 WHERE id = $4`
	_, err := db.ExecContext(ctx, stmt, r.Review, r.Rating, r.UpdatedAt, r.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteReviewByID deletes a review by its id (make sure we also delete the reference
// in book_reviews table)
func (r *Review) DeleteReviewByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM reviews WHERE id = $1`
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	stmt2 := `DELETE FROM book_reviews WHERE review_id = $1`
	_, err = db.ExecContext(ctx, stmt2, id)
	if err != nil {
		return err
	}

	return nil
}

// ReviewsByUserID returns all reviews by user id
func (r *Review) ReviewsByUserID(userID int) ([]*Review, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, review, rating, created_at, updated_at FROM reviews WHERE id IN (SELECT review_id FROM book_reviews WHERE user_id = $1) ORDER BY id`
	rows, err := db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*Review

	for rows.Next() {
		var review Review
		err := rows.Scan(&review.ID, &review.Review, &review.Rating, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, nil
}

// ReviewsByBookID returns all reviews by book id
func (r *Review) ReviewsByBookID(bookID int) ([]*Review, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, review, rating, created_at, updated_at FROM reviews WHERE id IN (SELECT review_id FROM book_reviews WHERE book_id = $1) ORDER BY id`
	rows, err := db.QueryContext(ctx, query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*Review

	for rows.Next() {
		var review Review
		err := rows.Scan(&review.ID, &review.Review, &review.Rating, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, nil
}
