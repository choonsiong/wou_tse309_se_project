package main

import "net/http"

func (app *application) AuthenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := app.models.Token.AuthenticateToken(r)
		if err != nil {
			app.errorLog.Println(err)
			payload := jsonResponse{
				Error:   true,
				Message: err.Error(),
			}

			_ = app.writeJSON(w, http.StatusUnauthorized, payload)
			return
		}

		next.ServeHTTP(w, r)
	})
}
