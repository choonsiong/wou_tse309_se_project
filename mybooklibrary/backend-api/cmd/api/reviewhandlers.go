package main

import "net/http"

// AllReviews handles API call to get all reviews
func (app *application) AllReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := app.models.Review.All()
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
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorLog.Println(err)
		_ = app.errorJSON(w, err)
		return
	}

	reviews, err := app.models.Review.ReviewsByBookID(requestPayload.ID)
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
