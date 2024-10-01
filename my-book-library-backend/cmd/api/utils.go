package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
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
	_, err = w.Write(out)
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

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(w, httpStatusCode, payload)
}
