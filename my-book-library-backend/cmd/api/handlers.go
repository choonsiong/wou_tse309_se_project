package main

import (
	"encoding/json"
	"net/http"
)

// jsonResponse type to structure JSON response
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

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		// user authentication failed
		app.errorLog.Println(err)

		payload.Error = true
		payload.Message = "invalid credentials"

		out, err := json.MarshalIndent(payload, "", "\t")
		if err != nil {
			app.errorLog.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write(out)

		return
	}

	// TODO: Authenticate user
	app.infoLog.Println(cred.Email, cred.Password)

	// send back JSON response
	payload.Error = false
	payload.Message = "success"

	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		app.errorLog.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(out)
}
