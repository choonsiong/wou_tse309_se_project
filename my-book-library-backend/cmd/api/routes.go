package main

import (
	"app-backend/internal/data/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"strconv"
	"time"
)

const (
	secret = "$2a$12$BqyqTZdjKcV6kYXJGqjukOKMRUEO5qHbSeC4457HwHWyk9MEKEHni"
)

func checkTestSecret(s string) bool {
	if s != secret {
		return false
	}

	return true
}

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
	mux.Post("/users/logout", app.Logout)

	mux.Route("/admin", func(r chi.Router) {
		r.Use(app.AuthenticateToken)

		r.Post("/test", func(w http.ResponseWriter, r *http.Request) {
			payload := jsonResponse{
				Error:   false,
				Message: "test",
			}

			_ = app.writeJSON(w, http.StatusOK, payload)
		})
	})

	mux.Get("/test/bcrypt", app.GenerateBcryptPassword)

	//mux.Get("/test/users/login", app.Login)

	mux.Get("/test/users/all", func(w http.ResponseWriter, r *http.Request) {
		secret := r.URL.Query().Get("secret")
		if !checkTestSecret(secret) {
			return
		}

		var u models.User

		users, err := u.GetAll()
		if err != nil {
			app.errorLog.Println(err)
		}

		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data: envelop{
				"users": users,
			},
		}

		_ = app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test/users/add", func(w http.ResponseWriter, r *http.Request) {
		secret := r.URL.Query().Get("secret")
		if !checkTestSecret(secret) {
			return
		}

		u := models.User{
			Email:     "alice@example.com",
			FirstName: "Alice",
			LastName:  "Smith",
			Password:  "password",
		}

		app.infoLog.Println("adding user")

		id, err := app.models.User.Insert(u)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err, http.StatusForbidden)
			return
		}

		app.infoLog.Println("user id added:", id)

		u2, _ := app.models.User.GetById(id)
		_ = app.writeJSON(w, http.StatusCreated, u2)
	})

	mux.Get("/test/tokens/generate", func(w http.ResponseWriter, r *http.Request) {
		secret := r.URL.Query().Get("secret")
		if !checkTestSecret(secret) {
			return
		}

		t, err := app.models.User.Token.GenerateToken(15, time.Minute*60)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err, http.StatusBadRequest)
			return
		}

		t.Email = "alice@example.com"
		t.CreatedAt = time.Now()
		t.UpdatedAt = time.Now()

		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data:    t,
		}

		_ = app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test/tokens/save", func(w http.ResponseWriter, r *http.Request) {
		secret := r.URL.Query().Get("secret")
		if !checkTestSecret(secret) {
			return
		}

		uidStr := r.URL.Query().Get("uid")
		if uidStr == "" {
			app.errorLog.Println("uidStr is empty")
			return
		}

		uid, err := strconv.Atoi(uidStr)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err, http.StatusBadRequest)
			return
		}

		u, err := app.models.User.GetById(uid)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err, http.StatusBadRequest)
			return
		}

		t, err := app.models.User.Token.GenerateToken(u.ID, time.Minute*60)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err, http.StatusBadRequest)
			return
		}

		u2, err := app.models.User.GetById(u.ID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err, http.StatusBadRequest)
			return
		}

		t.UserID = u2.ID
		t.CreatedAt = time.Now()
		t.UpdatedAt = time.Now()

		err = t.Insert(*t, *u2)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err, http.StatusBadRequest)
			return
		}

		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data:    t,
		}

		_ = app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test/tokens/validate", func(w http.ResponseWriter, r *http.Request) {
		secret := r.URL.Query().Get("secret")
		if !checkTestSecret(secret) {
			return
		}

		tokenStr := r.URL.Query().Get("token")
		isValid, err := app.models.Token.ValidToken(tokenStr)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err, http.StatusBadRequest)
			return
		}

		var payload jsonResponse
		payload.Error = false
		payload.Data = isValid

		_ = app.writeJSON(w, http.StatusOK, payload)
	})

	return mux
}
