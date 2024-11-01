package main

import (
	"app-backend/internal/data/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

// GetUserByID get user by the user id
func (app *application) GetUserByID(w http.ResponseWriter, r *http.Request) {
	uid, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	user, err := app.models.User.GetById(uid)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, user)
}

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

// DeleteUser handle the API for deleting user
func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	// First, get all books of the given user by the users_mybooks table
	userBooks, err := app.models.UserBook.GetAllBooksByUserID(requestPayload.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	// Then, delete all the records related to the given user id
	// in the users_mybooks table
	err = app.models.UserBook.DeleteByUserID(requestPayload.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	// Now, delete all the related references for each book owned by the given
	// user id
	for _, userBook := range userBooks {
		var err error

		err = app.models.Genre.DeleteForBookId(userBook.BookID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}

		err = app.models.Author.DeleteForBookId(userBook.BookID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}

		bookReviews, err := app.models.BookReview.GetByBookID(userBook.BookID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}

		bookReviewsUser, err := app.models.BookReview.GetByUserID(requestPayload.ID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}

		err = app.models.BookReview.DeleteByBookID(userBook.BookID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}

		err = app.models.BookReview.DeleteByUserID(requestPayload.ID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}

		for _, bookReview := range bookReviews {
			err := app.models.Review.DeleteReviewByID(bookReview.ReviewID)
			if err != nil {
				app.errorLog.Println(err)
				_ = app.errorJSON(w, err)
				return
			}
		}

		for _, bookReview := range bookReviewsUser {
			err := app.models.Review.DeleteReviewByID(bookReview.ReviewID)
			if err != nil {
				app.errorLog.Println(err)
				_ = app.errorJSON(w, err)
				return
			}
		}

		if userBook.UserID == requestPayload.ID {
			err = app.models.Book.DeleteByID(userBook.BookID)
			if err != nil {
				app.errorLog.Println(err)
				_ = app.errorJSON(w, err)
				return
			}
		}
	}

	bookReviewsUser, err := app.models.BookReview.GetByUserID(requestPayload.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	err = app.models.BookReview.DeleteByUserID(requestPayload.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	for _, bookReview := range bookReviewsUser {
		err := app.models.Review.DeleteReviewByID(bookReview.ReviewID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
	}

	// Finally delete the user from the database
	err = app.models.User.DeleteById(requestPayload.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
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
