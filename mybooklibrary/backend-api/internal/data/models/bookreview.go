package models

import (
	"context"
	"time"
)

type BookReview struct {
	ID        int       `json:"id"`
	ReviewID  int       `json:"review_id"`
	UserID    int       `json:"user_id"`
	BookID    int       `json:"book_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Insert inserts new record to the book_reviews table
func (b *BookReview) Insert(reviewId int, bookId int, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `INSERT INTO book_reviews (review_id, user_id, book_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var newID int

	err := db.QueryRowContext(ctx, stmt,
		reviewId,
		userId,
		bookId,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

// GetByID get row by id
func (b *BookReview) GetByID(id int) (*BookReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, review_id, user_id, book_id, created_at, updated_at FROM book_reviews WHERE id = $1`
	var bookReview BookReview

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(&bookReview.ID, &bookReview.ReviewID, &bookReview.UserID, &bookReview.BookID, &bookReview.CreatedAt, &bookReview.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &bookReview, nil
}

// GetByUserID returns rows by user id
func (b *BookReview) GetByUserID(id int) ([]*BookReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, review_id, user_id, book_id, created_at, updated_at FROM book_reviews WHERE user_id = $1`
	dbRows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var bookReviews []*BookReview

	for dbRows.Next() {
		var bookReview BookReview
		err := dbRows.Scan(
			&bookReview.ID,
			&bookReview.ReviewID,
			&bookReview.UserID,
			&bookReview.BookID,
			&bookReview.CreatedAt,
			&bookReview.UpdatedAt)
		if err != nil {
			return nil, err
		}

		bookReviews = append(bookReviews, &bookReview)
	}

	return bookReviews, nil
}

// GetByBookID returns rows by book id
func (b *BookReview) GetByBookID(id int) ([]*BookReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, review_id, user_id, book_id, created_at, updated_at FROM book_reviews WHERE book_id = $1`
	dbRows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var bookReviews []*BookReview

	for dbRows.Next() {
		var bookReview BookReview
		err := dbRows.Scan(
			&bookReview.ID,
			&bookReview.ReviewID,
			&bookReview.UserID,
			&bookReview.BookID,
			&bookReview.CreatedAt,
			&bookReview.UpdatedAt)
		if err != nil {
			return nil, err
		}

		bookReviews = append(bookReviews, &bookReview)
	}

	return bookReviews, nil
}

// DeleteByID delete row by id
func (b *BookReview) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM book_reviews WHERE id = $1`
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByReviewID delete all rows by review id
func (b *BookReview) DeleteByReviewID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM book_reviews WHERE review_id = $1`
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByBookID delete all rows by book id
func (b *BookReview) DeleteByBookID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM book_reviews WHERE book_id = $1`
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByUserID delete all rows by user id
func (b *BookReview) DeleteByUserID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM book_reviews WHERE user_id = $1`
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
