package main

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// envelop wraps data in JSON response
type envelop map[string]interface{}

// jsonResponse type to store a structure JSON response
type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
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

// Logout handler
func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &requestData)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, errors.New("error reading json from request data"))
		return
	}

	err = app.models.Token.DeleteByToken(requestData.Token)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, errors.New("failed to delete token"))
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "logged out",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
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
	//app.infoLog.Println("Authenticating user: ", cred.Email, cred.Password)

	// Search user by email
	u, err := app.models.User.GetByEmail(cred.Email)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, errors.New("invalid user"))
		return
	}

	// Validate user password
	isPasswordValid, err := u.PasswordMatches(cred.Password)
	if err != nil || !isPasswordValid {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, errors.New("invalid user"))
		return
	}

	// If we have a valid user, then generate a token
	t, err := app.models.Token.GenerateToken(u.ID, time.Hour*24)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, errors.New("failed to generate token"))
		return
	}

	// Save the token to the database
	err = app.models.Token.Insert(*t, *u)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, errors.New("failed to save token"))
		return
	}

	// Send back a JSON response if user successfully authenticated.
	payload = jsonResponse{
		Error:   false,
		Message: "user logged in",
		Data: envelop{
			"token": t,
			"user":  u,
		},
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
