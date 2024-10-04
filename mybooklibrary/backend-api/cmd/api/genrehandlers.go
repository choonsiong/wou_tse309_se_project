package main

import "net/http"

// AllGenres handles HTTP API call to get all genres
func (app *application) AllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.Genre.All()
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data: envelop{
			"genres": genres,
		},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
