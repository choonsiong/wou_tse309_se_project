package main

import (
	"app-backend/internal/data/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

// routes function handle application wide HTTP routes
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// The middlewares used in the application.
	mux.Use(middleware.Recoverer) // Make sure we recover from crash
	mux.Use(cors.Handler(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedMethods:   []string{"GET", "PATCH", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedOrigins:   []string{"https://*", "http://*"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300, // seconds
	})) // For cross-site server request

	mux.Post("/users/login", app.Login)
	
	mux.Get("/test/users/login", app.Login)
	mux.Get("/test/users/all", func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		users, err := u.GetAll()
		if err != nil {
			app.errorLog.Println(err)
		}

		_ = app.writeJSON(w, http.StatusOK, users)
	})

	return mux
}
