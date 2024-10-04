package main

import (
	"app-backend/internal/data/models"
	"encoding/base64"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/mozillazg/go-slugify"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var staticPath = "./static/"

// AllBooks is the HTTP handler for getting all books
func (app *application) AllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := app.models.Book.GetAll()
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data: envelop{
			"books": books,
		},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// GetBookByID is the HTTP handler to get a book by its id
func (app *application) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	book, err := app.models.Book.GetOneById(id)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data: envelop{
			"book": book,
		},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// DeleteBookByID is the HTTP handler to delete a book by its id
func (app *application) DeleteBook(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	err = app.models.Genre.DeleteForBookId(requestPayload.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	err = app.models.Author.DeleteForBookId(requestPayload.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	err = app.models.Book.DeleteByID(requestPayload.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) NewBook(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Authors         string `json:"authors"`
		BookCoverBase64 string `json:"book_cover"`
		Description     string `json:"description"`
		Genres          string `json:"genres"`
		PublisherName   string `json:"publisher_name"`
		PublicationYear int    `json:"publication_year"`
		Title           string `json:"title"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	//app.infoLog.Println(requestPayload.Title)
	//app.infoLog.Println(requestPayload.PublisherName)
	//app.infoLog.Println(requestPayload.PublicationYear)
	//app.infoLog.Println(requestPayload.Authors)
	//app.infoLog.Println(requestPayload.Genres)
	//app.infoLog.Println(requestPayload.Description)
	//app.infoLog.Println(requestPayload.BookCoverBase64)

	// Handle publisher
	// Get all existing publishers from the database
	publishers, err := app.models.Publisher.All()
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	foundPublisher := false
	publisherId := -1

	// Check whether the publisher given exists in the database
	for _, publisher := range publishers {
		if app.checkTwoStringsAreEqual(publisher.PublisherName, requestPayload.PublisherName) {
			foundPublisher = true
			publisherId = publisher.ID
			break
		}
	}

	// If the publisher given is not found
	if !foundPublisher {
		p := models.Publisher{
			PublisherName: app.normalizedString(requestPayload.PublisherName),
		}
		id, err := app.models.Publisher.Insert(p)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		// Set the new publisher id just inserted
		publisherId = id
	}

	// Get the new publisher from database
	publisher, err := app.models.Publisher.GetByID(publisherId)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	// Handle author
	// Get all authors from the database
	authors, err := app.models.Author.All()
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	var authorIds []int
	var authorNames []string

	authorsSplit := strings.Split(requestPayload.Authors, ",")

	// Loop through all the new authors given and compare with those in
	// the database
	for _, authorToCheck := range authorsSplit {
		found := false
		normalizedAuthor := app.normalizedString(authorToCheck)

		// Loop through all the authors in the database
		for _, authorInDB := range authors {
			if app.checkTwoStringsAreEqual(authorInDB.AuthorName, normalizedAuthor) {
				authorIds = append(authorIds, authorInDB.ID)
				authorNames = append(authorNames, app.normalizedString(authorInDB.AuthorName))
				found = true
			}
		}

		// If we cannot find the author name in the database
		if !found {
			a := models.Author{
				AuthorName: app.normalizedString(authorToCheck),
			}

			// Insert new author to database
			id, err := app.models.Author.Insert(a)
			if err != nil {
				app.errorLog.Println(err)
				_ = app.errorJSON(w, err)
				return
			}

			authorIds = append(authorIds, id)
			authorNames = append(authorNames, app.normalizedString(authorToCheck))
		}
	}

	// Now we use above authorIds slice to retrieve all authors again, and
	// this will be used in the book property
	var finalAuthors []models.Author

	for _, id := range authorIds {
		author, err := app.models.Author.GetByID(id)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		finalAuthors = append(finalAuthors, *author)
	}

	// Handle genres
	// Get all genres from the database
	genres, err := app.models.Genre.All()
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	var genreIds []int
	var genreNames []string

	genresSplit := strings.Split(requestPayload.Genres, ",")

	for _, genreToCheck := range genresSplit {
		found := false
		normalizedGenre := app.normalizedString(genreToCheck)

		for _, genreInDB := range genres {
			if app.checkTwoStringsAreEqual(genreInDB.GenreName, normalizedGenre) {
				genreIds = append(genreIds, genreInDB.ID)
				genreNames = append(genreNames, app.normalizedString(genreInDB.GenreName))
				found = true
				break
			}
		}

		if !found {
			g := models.Genre{
				GenreName: app.normalizedString(genreToCheck),
			}

			id, err := app.models.Genre.Insert(g)
			if err != nil {
				app.errorLog.Println(err)
				_ = app.errorJSON(w, err)
				return
			}

			genreIds = append(genreIds, id)
			genreNames = append(genreNames, app.normalizedString(genreToCheck))
		}
	}

	var finalGenres []models.Genre

	for _, id := range genreIds {
		genre, err := app.models.Genre.GetByID(id)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		finalGenres = append(finalGenres, *genre)
	}

	// Prepare to process new book

	b := models.Book{
		Title:           requestPayload.Title,
		PublisherID:     publisherId,
		PublicationYear: requestPayload.PublicationYear,
		Description:     requestPayload.Description,
		Slug:            slugify.Slugify(requestPayload.Title),
		Publisher:       *publisher,
		Authors:         finalAuthors,
		AuthorIDs:       authorIds,
		Genres:          finalGenres,
		GenreIDs:        genreIds,
	}

	// Process book cover image
	if len(requestPayload.BookCoverBase64) > 0 {
		decoded, err := base64.StdEncoding.DecodeString(requestPayload.BookCoverBase64)
		if err != nil {
			_ = app.errorJSON(w, err)
			return
		}

		// Write image to /static/img
		if err := os.WriteFile(fmt.Sprintf("%s/img/%s.jpg", staticPath, b.Slug), decoded, 0644); err != nil {
			_ = app.errorJSON(w, err)
			return
		}
	}

	b.Details()

	bookId, err := app.models.Book.Insert(b)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data:    bookId,
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// EditBook is the HTTP handler to edit a book
func (app *application) EditBook(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID              int    `json:"id"`
		Authors         string `json:"authors"`
		BookCoverBase64 string `json:"book_cover"`
		Description     string `json:"description"`
		Genres          string `json:"genres"`
		PublisherName   string `json:"publisher_name"`
		PublicationYear int    `json:"publication_year"`
		Slug            string `json:"slug"`
		Title           string `json:"title"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	//app.infoLog.Println(requestPayload.ID)
	//app.infoLog.Println(requestPayload.Title)
	//app.infoLog.Println(requestPayload.PublisherName)
	//app.infoLog.Println(requestPayload.PublicationYear)
	//app.infoLog.Println(requestPayload.Authors)
	//app.infoLog.Println(requestPayload.Genres)
	//app.infoLog.Println(requestPayload.Description)
	//app.infoLog.Println(requestPayload.Slug)
	//app.infoLog.Println(requestPayload.BookCoverBase64)

	// Handle publisher
	// Get all existing publishers from the database
	publishers, err := app.models.Publisher.All()
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	foundPublisher := false
	publisherId := -1

	// Check whether the publisher given exists in the database
	for _, publisher := range publishers {
		if app.checkTwoStringsAreEqual(publisher.PublisherName, requestPayload.PublisherName) {
			foundPublisher = true
			publisherId = publisher.ID
			break
		}
	}

	// If the publisher given is not found
	if !foundPublisher {
		p := models.Publisher{
			PublisherName: app.normalizedString(requestPayload.PublisherName),
		}
		id, err := app.models.Publisher.Insert(p)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		// Set the new publisher id just inserted
		publisherId = id
	}

	// Get the new publisher from database
	publisher, err := app.models.Publisher.GetByID(publisherId)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	// Handle author
	// Get all authors from the database
	authors, err := app.models.Author.All()
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	var authorIds []int
	var authorNames []string

	authorsSplit := strings.Split(requestPayload.Authors, ",")

	// Loop through all the new authors given and compare with those in
	// the database
	for _, authorToCheck := range authorsSplit {
		found := false
		normalizedAuthor := app.normalizedString(authorToCheck)

		// Loop through all the authors in the database
		for _, authorInDB := range authors {
			if app.checkTwoStringsAreEqual(authorInDB.AuthorName, normalizedAuthor) {
				authorIds = append(authorIds, authorInDB.ID)
				authorNames = append(authorNames, app.normalizedString(authorInDB.AuthorName))
				found = true
			}
		}

		// If we cannot find the author name in the database
		if !found {
			a := models.Author{
				AuthorName: app.normalizedString(authorToCheck),
			}

			// Insert new author to database
			id, err := app.models.Author.Insert(a)
			if err != nil {
				app.errorLog.Println(err)
				_ = app.errorJSON(w, err)
				return
			}

			authorIds = append(authorIds, id)
			authorNames = append(authorNames, app.normalizedString(authorToCheck))
		}
	}

	// Now we use above authorIds slice to retrieve all authors again, and
	// this will be used in the book property
	var finalAuthors []models.Author

	for _, id := range authorIds {
		author, err := app.models.Author.GetByID(id)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		finalAuthors = append(finalAuthors, *author)
	}

	// Handle genres
	// Get all genres from the database
	genres, err := app.models.Genre.All()
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	var genreIds []int
	var genreNames []string

	genresSplit := strings.Split(requestPayload.Genres, ",")

	for _, genreToCheck := range genresSplit {
		found := false
		normalizedGenre := app.normalizedString(genreToCheck)

		for _, genreInDB := range genres {
			if app.checkTwoStringsAreEqual(genreInDB.GenreName, normalizedGenre) {
				genreIds = append(genreIds, genreInDB.ID)
				genreNames = append(genreNames, app.normalizedString(genreInDB.GenreName))
				found = true
				break
			}
		}

		if !found {
			g := models.Genre{
				GenreName: app.normalizedString(genreToCheck),
			}

			id, err := app.models.Genre.Insert(g)
			if err != nil {
				app.errorLog.Println(err)
				_ = app.errorJSON(w, err)
				return
			}

			genreIds = append(genreIds, id)
			genreNames = append(genreNames, app.normalizedString(genreToCheck))
		}
	}

	var finalGenres []models.Genre

	for _, id := range genreIds {
		genre, err := app.models.Genre.GetByID(id)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		finalGenres = append(finalGenres, *genre)
	}

	// Prepare to process new book

	b := models.Book{
		ID:              requestPayload.ID,
		Title:           requestPayload.Title,
		PublisherID:     publisherId,
		PublicationYear: requestPayload.PublicationYear,
		Description:     requestPayload.Description,
		Slug:            slugify.Slugify(requestPayload.Title),
		Publisher:       *publisher,
		Authors:         finalAuthors,
		AuthorIDs:       authorIds,
		Genres:          finalGenres,
		GenreIDs:        genreIds,
	}

	// Process book cover image
	if len(requestPayload.BookCoverBase64) > 0 {
		decoded, err := base64.StdEncoding.DecodeString(requestPayload.BookCoverBase64)
		if err != nil {
			_ = app.errorJSON(w, err)
			return
		}

		// Write image to /static/img
		if err := os.WriteFile(fmt.Sprintf("%s/img/%s.jpg", staticPath, b.Slug), decoded, 0644); err != nil {
			_ = app.errorJSON(w, err)
			return
		}
	}

	b.Details()

	err = b.Update()
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
