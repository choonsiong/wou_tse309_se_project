package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// routes function handle HTTP routes
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// middlewares
	mux.Use(middleware.Recoverer)

	mux.Get("/users/login", app.Login)
	mux.Post("/users/login", app.Login)

	return mux
}
