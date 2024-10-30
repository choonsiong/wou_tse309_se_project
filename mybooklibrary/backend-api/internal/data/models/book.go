package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mozillazg/go-slugify"
	"log"
	"time"
)

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	PublisherID     int       `json:"publisher_id"`
	PublicationYear int       `json:"publication_year"`
	Description     string    `json:"description"`
	ISBN            string    `json:"isbn"`
	Slug            string    `json:"slug"`
	Publisher       Publisher `json:"publisher"`
	Authors         []Author  `json:"authors"`
	AuthorIDs       []int     `json:"author_ids,omitempty"`
	Genres          []Genre   `json:"genres"`
	GenreIDs        []int     `json:"genre_ids,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (b *Book) Details() {
	log.Println("ID:", b.ID)
	log.Println("Title:", b.Title)
	log.Println("Publisher:", b.PublisherID)
	log.Println("Publication Year:", b.PublicationYear)
	log.Println("Description:", b.Description)
	log.Println("ISBN:", b.ISBN)
	log.Println("Slug:", b.Slug)
	log.Println("Publisher:", b.Publisher.PublisherName)
	for i, author := range b.Authors {
		log.Println("Index:", i, ", Author:", author.AuthorName, ", Author ID: ", author.ID)
	}
	for i, authorId := range b.AuthorIDs {
		log.Println("Index:", i, ", AuthorID:", authorId)
	}
	for i, genre := range b.Genres {
		log.Println("Index:", i, ", Genre:", genre.GenreName, ", Genre ID: ", genre.ID)
	}
	for i, genreId := range b.GenreIDs {
		log.Println("Index:", i, ", GenreID:", genreId)
	}
	log.Println("CreatedAt:", b.CreatedAt)
	log.Println("UpdatedAt:", b.UpdatedAt)
}

// GetOneById returns one book by its id
func (b *Book) GetOneById(id int) (*Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT b.id, b.title, b.publisher_id, b.publication_year, b.description, b.isbn, b.slug, b.created_at, b.updated_at, a.id, a.publisher_name, a.created_at, a.updated_at FROM mybooks b LEFT JOIN publishers a ON (b.publisher_id = a.id) WHERE b.id = $1`

	row := db.QueryRowContext(ctx, query, id)

	var book Book

	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.PublisherID,
		&book.PublicationYear,
		&book.Description,
		&book.ISBN,
		&book.Slug,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.Publisher.ID,
		&book.Publisher.PublisherName,
		&book.Publisher.CreatedAt,
		&book.Publisher.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Get genres
	genres, ids, err := b.genresForBook(book.ID)
	if err != nil {
		return nil, err
	}
	book.Genres = genres
	book.GenreIDs = ids

	// Get authors
	authors, ids, err := b.authorsForBook(book.ID)
	if err != nil {
		return nil, err
	}
	book.Authors = authors
	book.AuthorIDs = ids

	return &book, nil
}

// GetOneBySlug returns one book by slug
func (b *Book) GetOneBySlug(slug string) (*Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT b.id, b.title, b.publisher_id, b.publication_year, b.description, b.isbn, b.slug, b.created_at, b.updated_at, a.id, a.publisher_name, a.created_at, a.updated_at FROM mybooks b LEFT JOIN publishers a ON (b.publisher_id = a.id) WHERE b.slug = $1`

	row := db.QueryRowContext(ctx, query, slug)

	var book Book

	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.PublisherID,
		&book.PublicationYear,
		&book.Description,
		&book.ISBN,
		&book.Slug,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.Publisher.ID,
		&book.Publisher.PublisherName,
		&book.Publisher.CreatedAt,
		&book.Publisher.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// get genres
	genres, ids, err := b.genresForBook(book.ID)
	if err != nil {
		return nil, err
	}
	book.Genres = genres
	book.GenreIDs = ids

	// Get authors
	authors, ids, err := b.authorsForBook(book.ID)
	if err != nil {
		return nil, err
	}
	book.Authors = authors
	book.AuthorIDs = ids

	return &book, nil
}

// GetAll returns a slice of all books
func (b *Book) GetAll() ([]*Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT b.id, b.title, b.publisher_id, b.publication_year, b.description, b.isbn, b.slug, b.created_at, b.updated_at, a.id, a.publisher_name, a.created_at, a.updated_at FROM mybooks b LEFT JOIN publishers a ON (b.publisher_id = a.id) ORDER BY b.created_at DESC`

	var books []*Book

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.PublisherID,
			&book.PublicationYear,
			&book.Description,
			&book.ISBN,
			&book.Slug,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.Publisher.ID,
			&book.Publisher.PublisherName,
			&book.Publisher.CreatedAt,
			&book.Publisher.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Get genres
		genres, ids, err := b.genresForBook(book.ID)
		if err != nil {
			return nil, err
		}
		book.Genres = genres
		book.GenreIDs = ids

		// Get authors
		authors, ids, err := b.authorsForBook(book.ID)
		if err != nil {
			return nil, err
		}
		book.Authors = authors
		book.AuthorIDs = ids

		books = append(books, &book)
	}

	return books, nil
}

// GetAllPaginated returns a slice of all books, paginated by limit and offset
func (b *Book) GetAllPaginated(page, pageSize int) ([]*Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	limit := pageSize
	offset := (page - 1) * pageSize

	query := `SELECT b.id, b.title, b.publisher_id, b.publication_year, b.description, b.isbn, b.slug, b.created_at, b.updated_at, a.id, a.publisher_name, a.created_at, a.updated_at FROM mybooks b LEFT JOIN publishers a ON (b.publisher_id = a.id) ORDER BY b.title LIMIT $1 OFFSET $2`

	var books []*Book

	rows, err := db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.PublisherID,
			&book.PublicationYear,
			&book.Description,
			&book.ISBN,
			&book.Slug,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.Publisher.ID,
			&book.Publisher.PublisherName,
			&book.Publisher.CreatedAt,
			&book.Publisher.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Get genres
		genres, ids, err := b.genresForBook(book.ID)
		if err != nil {
			return nil, err
		}
		book.Genres = genres
		book.GenreIDs = ids

		// Get authors
		authors, ids, err := b.authorsForBook(book.ID)
		if err != nil {
			return nil, err
		}
		book.Authors = authors
		book.AuthorIDs = ids

		books = append(books, &book)
	}

	return books, nil
}

// Insert saves one book to the database
func (b *Book) Insert(book Book) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `INSERT INTO mybooks (title, publisher_id, publication_year, description, isbn, slug, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	var newID int
	err := db.QueryRowContext(ctx, stmt,
		book.Title,
		book.PublisherID,
		book.PublicationYear,
		book.Description,
		book.ISBN,
		slugify.Slugify(book.Title),
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// Update genres using genre ids
	if len(book.GenreIDs) > 0 {
		stmt = `DELETE FROM mybooks_genres WHERE book_id = $1`
		_, err := db.ExecContext(ctx, stmt, book.ID)
		if err != nil {
			return newID, fmt.Errorf("book updated, but genres not: %s", err.Error())
		}

		// Add new genres
		for _, x := range book.GenreIDs {
			stmt = `INSERT INTO mybooks_genres (book_id, genre_id, created_at, updated_at) VALUES ($1, $2, $3, $4)`
			_, err = db.ExecContext(ctx, stmt, newID, x, time.Now(), time.Now())
			if err != nil {
				return newID, fmt.Errorf("book updated, but genres not: %s", err.Error())
			}
		}
	}

	// Update authors using author ids
	if len(book.AuthorIDs) > 0 {
		stmt = `DELETE FROM mybooks_authors WHERE book_id = $1`
		_, err := db.ExecContext(ctx, stmt, book.ID)
		if err != nil {
			return newID, fmt.Errorf("book updated, but authors not: %s", err.Error())
		}

		// Add new authors
		for _, x := range book.AuthorIDs {
			stmt = `INSERT INTO mybooks_authors (book_id, author_id, created_at, updated_at) VALUES ($1, $2, $3, $4)`
			_, err = db.ExecContext(ctx, stmt, newID, x, time.Now(), time.Now())
			if err != nil {
				return newID, fmt.Errorf("book updated, but authors not: %s", err.Error())
			}
		}
	}

	return newID, nil
}

// Update updates one book in the database
func (b *Book) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `UPDATE mybooks SET
		title = $1,
		publisher_id = $2,
		publication_year = $3,
		description = $4,
		isbn = $5,
        slug = $6,
		updated_at = $7
		where id = $8`

	_, err := db.ExecContext(
		ctx,
		stmt,
		b.Title,
		b.PublisherID,
		b.PublicationYear,
		b.Description,
		b.ISBN,
		slugify.Slugify(b.Title),
		time.Now(),
		b.ID)
	if err != nil {
		return err
	}

	// Update genres using genre ids
	if len(b.GenreIDs) > 0 {
		stmt = `DELETE from mybooks_genres WHERE book_id = $1`
		_, err := db.ExecContext(ctx, stmt, b.ID)
		if err != nil {
			return fmt.Errorf("book updated, but genres not: %s", err.Error())
		}

		// Add new genres
		for _, x := range b.GenreIDs {
			stmt = `INSERT INTO mybooks_genres (book_id, genre_id, created_at, updated_at) VALUES ($1, $2, $3, $4)`
			_, err = db.ExecContext(ctx, stmt, b.ID, x, time.Now(), time.Now())
			if err != nil {
				return fmt.Errorf("book updated, but genres not: %s", err.Error())
			}
		}
	}

	// Update authors using author ids
	if len(b.AuthorIDs) > 0 {
		stmt = `DELETE from mybooks_authors WHERE book_id = $1`
		_, err := db.ExecContext(ctx, stmt, b.ID)
		if err != nil {
			return fmt.Errorf("book updated, but authors not: %s", err.Error())
		}

		// Add new authors
		for _, x := range b.AuthorIDs {
			stmt = `INSERT INTO mybooks_authors (book_id, author_id, created_at, updated_at) VALUES ($1, $2, $3, $4)`
			_, err = db.ExecContext(ctx, stmt, b.ID, x, time.Now(), time.Now())
			if err != nil {
				return fmt.Errorf("book updated, but authors not: %s", err.Error())
			}
		}
	}

	return nil
}

// DeleteByID deletes a book by id
func (b *Book) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM mybooks WHERE id = $1`
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllByUserID get all books for the given user id.
func (b *Book) GetAllByUserID(id int) ([]*Book, []int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, title, publisher_id, publication_year, description, isbn, slug, created_at, updated_at FROM mybooks WHERE id IN (SELECT book_id FROM users_mybooks WHERE user_id = $1) ORDER BY id DESC`
	dbRows, err := db.QueryContext(ctx, query, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, err
	}
	defer dbRows.Close()

	var books []*Book
	var bookIds []int

	for dbRows.Next() {
		var book Book

		err = dbRows.Scan(
			&book.ID,
			&book.Title,
			&book.PublisherID,
			&book.PublicationYear,
			&book.Description,
			&book.ISBN,
			&book.Slug,
			&book.CreatedAt,
			&book.UpdatedAt)
		if err != nil {
			return nil, nil, err
		}

		// Get publisher
		publisher, err := b.publisherForBook(book)
		if err != nil {
			return nil, nil, err
		}
		book.Publisher = *publisher

		// Get genres
		genres, ids, err := b.genresForBook(book.ID)
		if err != nil {
			return nil, nil, err
		}
		book.Genres = genres
		book.GenreIDs = ids

		// Get authors
		authors, ids, err := b.authorsForBook(book.ID)
		if err != nil {
			return nil, nil, err
		}
		book.Authors = authors
		book.AuthorIDs = ids

		books = append(books, &book)
		bookIds = append(bookIds, book.ID)
	}

	return books, bookIds, nil
}

// publisherForBook returns publisher for the given book
func (b *Book) publisherForBook(book Book) (*Publisher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, publisher_name, created_at, updated_at FROM publishers WHERE id = $1`
	row := db.QueryRowContext(ctx, query, book.PublisherID)

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

// genresForBook returns all genres for a given book id
func (b *Book) genresForBook(id int) ([]Genre, []int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	// Get genres
	var genres []Genre
	var genreIDs []int
	genreQuery := `SELECT id, genre_name, created_at, updated_at FROM genres WHERE id IN (SELECT genre_id FROM mybooks_genres WHERE book_id = $1) ORDER BY genre_name`

	gRows, err := db.QueryContext(ctx, genreQuery, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, err
	}
	defer gRows.Close()

	var genre Genre
	for gRows.Next() {
		err = gRows.Scan(
			&genre.ID,
			&genre.GenreName,
			&genre.CreatedAt,
			&genre.UpdatedAt)
		if err != nil {
			return nil, nil, err
		}
		genres = append(genres, genre)
		genreIDs = append(genreIDs, genre.ID)
	}

	return genres, genreIDs, nil
}

// authorsForBook returns all authors for a given book id
func (b *Book) authorsForBook(id int) ([]Author, []int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	// Get authors
	var authors []Author
	var authorIDs []int
	authorQuery := `SELECT id, author_name, created_at, updated_at FROM authors WHERE id IN (SELECT author_id FROM mybooks_authors WHERE book_id = $1) ORDER BY author_name`

	aRows, err := db.QueryContext(ctx, authorQuery, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, err
	}
	defer aRows.Close()

	var author Author
	for aRows.Next() {

		err = aRows.Scan(
			&author.ID,
			&author.AuthorName,
			&author.CreatedAt,
			&author.UpdatedAt)
		if err != nil {
			return nil, nil, err
		}
		authors = append(authors, author)
		authorIDs = append(authorIDs, author.ID)
	}

	return authors, authorIDs, nil
}
