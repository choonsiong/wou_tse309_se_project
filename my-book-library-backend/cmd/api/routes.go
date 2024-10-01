package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// routes function handle application HTTP routes
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// The middlewares used in the application.
	mux.Use(middleware.Recoverer) // recover from crash

	mux.Get("/users/login", app.Login)
	mux.Post("/users/login", app.Login)

	return mux
}
