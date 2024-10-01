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

	//err := json.NewDecoder(r.Body).Decode(&cred)
	//if err != nil {
	//	// User authentication failed.
	//	app.errorLog.Println("Failed to authenticate user: ", err)
	//
	//	payload.Error = true
	//	payload.Message = "invalid credentials"
	//
	//	out, err := json.MarshalIndent(payload, "", "\t")
	//	if err != nil {
	//		app.errorLog.Println(err)
	//	}
	//
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusOK)
	//
	//	_, _ = w.Write(out)
	//
	//	return
	//}

	err := app.readJSON(w, r, &cred)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "failed to read JSON data from request body"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	// TODO: Authenticate user
	app.infoLog.Println("Authenticating user: ", cred.Email, cred.Password)

	// Send back a JSON response if user successfully authenticated.
	payload.Error = false
	payload.Message = "user authenticated successfully"

	//out, err := json.MarshalIndent(payload, "", "\t")
	//if err != nil {
	//	app.errorLog.Println(err)
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//
	//_, _ = w.Write(out)

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
