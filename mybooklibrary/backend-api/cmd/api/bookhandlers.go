package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// AllBooks handle HTTP API calls to get all books
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

// GetBookByID handle HTTP API calls to get one book by id
func (app *application) GetBookByID(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	book, err := app.models.Book.GetOneById(bookId)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	app.infoLog.Println(book.Publisher.PublisherName)
	app.infoLog.Println(len(book.Genres))

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data: envelop{
			"book": book,
		},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
