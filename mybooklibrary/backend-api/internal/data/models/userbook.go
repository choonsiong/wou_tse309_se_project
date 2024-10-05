package models

import (
	"context"
	"time"
)

type UserBook struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	BookID    int       `json:"book_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Insert inserts new record to the database
func (r *UserBook) Insert(userID int, bookID int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `INSERT INTO users_mybooks (user_id, book_id, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

	var newID int

	err := db.QueryRowContext(ctx, stmt,
		userID,
		bookID,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

// DeleteByID deletes all records by the given id
func (r *UserBook) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM users_mybooks WHERE id = $1`
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByUserID deletes all records by the given user id
func (r *UserBook) DeleteByUserID(userID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM users_mybooks WHERE user_id = $1`
	_, err := db.ExecContext(ctx, stmt, userID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByBookID deletes all records by the given book id
func (r *UserBook) DeleteByBookID(bookID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM users_mybooks WHERE book_id = $1`
	_, err := db.ExecContext(ctx, stmt, bookID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllBooksByUserID gets all books (i.e., the book ids) for the given user id
func (r *UserBook) GetAllBooksByUserID(id int) ([]*UserBook, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, user_id, book_id, created_at, updated_at FROM users_mybooks WHERE user_id = $1`

	dbRows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var userBooks []*UserBook

	for dbRows.Next() {
		var userBook UserBook
		err := dbRows.Scan(
			&userBook.ID,
			&userBook.UserID,
			&userBook.BookID,
			&userBook.CreatedAt,
			&userBook.UpdatedAt)
		if err != nil {
			return nil, err
		}

		userBooks = append(userBooks, &userBook)
	}

	return userBooks, nil
}

// GetAllUsersByBookID gets all users (i.e., the user id) for the given book id
func (r *UserBook) GetAllUsersByBookID(id int) ([]*UserBook, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, user_id, book_id, created_at, updated_at FROM users_mybooks WHERE book_id = $1`

	dbRows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var userBooks []*UserBook

	for dbRows.Next() {
		var userBook UserBook
		err := dbRows.Scan(
			&userBook.ID,
			&userBook.UserID,
			&userBook.BookID,
			&userBook.CreatedAt,
			&userBook.UpdatedAt)
		if err != nil {
			return nil, err
		}

		userBooks = append(userBooks, &userBook)
	}

	return userBooks, nil
}
