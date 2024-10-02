package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// readJSON is a helper function to read JSON data from HTTP request body.
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxReadBytes := 1048576 // 1MB

	// Only read 1MB of data
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxReadBytes))

	// Create a new decoder from request body
	d := json.NewDecoder(r.Body)
	err := d.Decode(data) // decode into data
	if err != nil {
		return err
	}

	err = d.Decode(&struct{}{}) // make sure we only received a single JSON value
	if err != io.EOF {
		return errors.New("request body must have only a single JSON value")
	}

	return nil
}

// writeJSON is a helper function to write JSON data into HTTP response body. httpHeaders parameter accepts zero or one value.
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, httpHeaders ...http.Header) error {
	var output []byte

	if app.cfg.ENV == "dev" {
		out, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			return err
		}
		output = out
	} else {
		out, err := json.Marshal(data)
		if err != nil {
			return err
		}
		output = out
	}

	// If we have additional HTTP headers
	if len(httpHeaders) > 0 {
		for k, v := range httpHeaders[0] { // in fact, we only accept a single http.Header value
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Write the JSON to response body
	_, err := w.Write(output)
	if err != nil {
		return err
	}

	return nil
}

// errorJSON is a helper function to write JSON data into HTTP response body for communicating error message. httpStatusCodes accepts zero or one value.
func (app *application) errorJSON(w http.ResponseWriter, err error, httpStatusCodes ...int) error {
	httpStatusCode := http.StatusBadRequest

	if len(httpStatusCodes) > 0 {
		httpStatusCode = httpStatusCodes[0]
	}

	var customErr error

	switch {
	case strings.Contains(err.Error(), "SQLSTATE 23505"):
		customErr = errors.New("duplicate value violates unique constraint")
		httpStatusCode = http.StatusForbidden
	case strings.Contains(err.Error(), "SQLSTATE 22001"):
		customErr = errors.New("the value you are tring to insert is too large")
		httpStatusCode = http.StatusForbidden
	case strings.Contains(err.Error(), "SQLSTATE 23503"):
		customErr = errors.New("foreign key violation")
		httpStatusCode = http.StatusForbidden
	default:
		customErr = err
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = customErr.Error()

	return app.writeJSON(w, httpStatusCode, payload)
}

// checkTestSecret check secret string for protected routes
func (app *application) checkTestSecret(s string) bool {
	if s != app.cfg.Secret {
		return false
	}

	return true
}
