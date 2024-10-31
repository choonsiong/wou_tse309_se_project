package main

import (
	"app-backend/internal/data/models"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

type ResponsePayload struct {
	ID        int       `json:"id"`
	Review    string    `json:"review"`
	Rating    int       `json:"rating"`
	Bookname  string    `json:"book_name"`
	Username  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AllReviews handles API call to get all reviews
func (app *application) AllReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := app.models.Review.All()
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	var results []ResponsePayload

	for _, review := range reviews {
		r, err := app.models.BookReview.GetByReviewID(review.ID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		b, err := app.models.Book.GetOneById(r.BookID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		u, err := app.models.User.GetById(r.UserID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		var p ResponsePayload
		p.ID = review.ID
		p.Review = review.Review
		p.Rating = review.Rating
		p.CreatedAt = review.CreatedAt
		p.UpdatedAt = review.UpdatedAt
		p.Bookname = b.Title
		p.Username = u.FirstName + " " + u.LastName
		results = append(results, p)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data: envelop{
			"results": results,
		},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// IsReviewExistsForSameUserAndBook handles API call to check whether review exists for same user id and book id
func (app *application) IsReviewExistsForSameUserAndBook(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		BookId int `json:"book_id"`
		UserId int `json:"user_id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	review, err := app.models.BookReview.GetOneByUserIdAndBookId(requestPayload.UserId, requestPayload.BookId)
	if err != nil && err != sql.ErrNoRows {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	var isReviewExists bool

	if review == nil {
		isReviewExists = false
	} else {
		isReviewExists = true
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data: envelop{
			"exists": isReviewExists,
		},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// NewReview handles API call to inserts a new review
func (app *application) NewReview(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		BookId int    `json:"book_id"`
		UserId int    `json:"user_id"`
		Review string `json:"review"`
		Rating int    `json:"rating"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	review := &models.Review{
		Review: requestPayload.Review,
		Rating: requestPayload.Rating,
	}

	_, _, err = app.models.Review.Insert(review.Review, review.Rating, requestPayload.UserId, requestPayload.BookId)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// EditReview handles API call to update review
func (app *application) EditReview(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Id int `json:"id"`
		//BookId int    `json:"book_id"`
		//UserId int    `json:"user_id"`
		Review string `json:"review"`
		Rating int    `json:"rating"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	review := &models.Review{
		ID:     requestPayload.Id,
		Review: requestPayload.Review,
		Rating: requestPayload.Rating,
	}

	err = review.Update()
	if err != nil {
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) DeleteReviewById(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	err = app.models.BookReview.DeleteByReviewID(requestPayload.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	err = app.models.Review.DeleteReviewByID(requestPayload.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// AllReviewsByUserID handles API call to get all reviews according to user id
func (app *application) AllReviewsByUserID(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	reviews, err := app.models.Review.ReviewsByUserID(requestPayload.ID)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data: envelop{
			"reviews": reviews,
		},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// AllReviewsByBookID handles API call to get all reviews according to book id
func (app *application) AllReviewsByBookID(w http.ResponseWriter, r *http.Request) {
	//var requestPayload struct {
	//	ID int `json:"id"`
	//}
	//
	//err := app.readJSON(w, r, &requestPayload)
	//if err != nil {
	//	app.errorLog.Println(err)
	//	_ = app.errorJSON(w, err)
	//	return
	//}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	reviews, err := app.models.Review.ReviewsByBookID(id)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	var results []ResponsePayload

	for _, review := range reviews {
		r, err := app.models.BookReview.GetByReviewID(review.ID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		b, err := app.models.Book.GetOneById(r.BookID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		u, err := app.models.User.GetById(r.UserID)
		if err != nil {
			app.errorLog.Println(err)
			_ = app.errorJSON(w, err)
			return
		}
		var p ResponsePayload
		p.ID = review.ID
		p.Review = review.Review
		p.Rating = review.Rating
		p.CreatedAt = review.CreatedAt
		p.UpdatedAt = review.UpdatedAt
		p.Bookname = b.Title
		p.Username = u.FirstName + " " + u.LastName
		results = append(results, p)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data: envelop{
			"results": results,
		},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
