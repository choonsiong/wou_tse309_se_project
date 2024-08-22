package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// application type store application wide shared data
type application struct {
	cfg      config
	errorLog *log.Logger
	infoLog  *log.Logger
}

// start the web server
func (app *application) serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}
		payload.Status = true
		payload.Message = "Hello World"
		out, err := json.MarshalIndent(payload, "", "\t")
		if err != nil {
			app.errorLog.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(out)
	})

	app.infoLog.Println("Listening on port ", app.cfg.port)

	return http.ListenAndServe(fmt.Sprintf(":%d", app.cfg.port), nil)
}
