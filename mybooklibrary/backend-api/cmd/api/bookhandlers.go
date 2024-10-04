package main

import "net/http"

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
