package main

import (
	"net/http"
)

// jsonResponse type to store a structure JSON response
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Login handler
func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credential struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var cred credential
	var payload jsonResponse

	err := app.readJSON(w, r, &cred)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "failed to read JSON data from request body"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
		return
	}

	// TODO: Authenticate user
	app.infoLog.Println("Authenticating user: ", cred.Email, cred.Password)

	// Send back a JSON response if user successfully authenticated.
	payload.Error = false
	payload.Message = "user authenticated successfully"

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
