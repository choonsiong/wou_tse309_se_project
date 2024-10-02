package main

import (
	"app-backend/internal/data/models"
	"net/http"
)

// GetAllUsers get all users from database
func (app *application) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		secret := r.URL.Query().Get("secret")
		if !app.checkTestSecret(secret) {
			app.errorLog.Println("invalid secret")
			return
		}
	}

	var u models.User

	users, err := u.GetAll()
	if err != nil {
		app.errorLog.Println(err)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data: envelop{
			"users": users,
		},
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
