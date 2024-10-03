package main

import (
	"app-backend/internal/data/models"
	"net/http"
	"time"
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
		return
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

// NewUser handle the API for creating a new user
func (app *application) NewUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := app.readJSON(w, r, &u)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	u.Description()

	if _, err := app.models.User.Insert(u); err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "user saved",
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)
}

// EditUser handle the API for updating an existing user
func (app *application) EditUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := app.readJSON(w, r, &u)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	//u.Description()

	// Make sure we have an existing user to update
	u2, err := app.models.User.GetById(u.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	u2.FirstName = u.FirstName
	u2.LastName = u.LastName
	u2.Email = u.Email
	u2.Active = u.Active
	u2.IsAdmin = u.IsAdmin
	u2.UpdatedAt = time.Now()

	u2.Description()

	if err := u2.Update(); err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	// If user want to update password
	if u.Password != "" {
		err := u.ResetPassword(u.Password)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
	}

	payload := jsonResponse{
		Error:   false,
		Message: "user updated",
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)
}
