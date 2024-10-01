package main

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// jsonResponse type to store a structure JSON response
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// GenerateBcryptPassword generates bcrypt password according to user provided query string
// http://localhost:9009/test/bcrypt?p=abc
func (app *application) GenerateBcryptPassword(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("p")
	bs := []byte(q)

	// Hashing the password with a cost of 12
	hashedPassword, err := bcrypt.GenerateFromPassword(bs, 12)
	if err != nil {
		app.errorLog.Println(err)
	}

	app.infoLog.Println("Generated bcrypt password:", string(hashedPassword))

	var payload jsonResponse
	payload.Error = false
	payload.Message = string(hashedPassword)

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
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
