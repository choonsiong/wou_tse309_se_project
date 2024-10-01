package main

import (
	"app-backend/internal/driver"
	"fmt"
	"log"
	"net/http"
)

// application type store application wide shared data
type application struct {
	cfg      config
	db       *driver.DB
	errorLog *log.Logger
	infoLog  *log.Logger
}

// start the web server
func (app *application) serve() error {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	var payload struct {
	//		Status  bool   `json:"status"`
	//		Message string `json:"message"`
	//	}
	//	payload.Status = true
	//	payload.Message = "Hello World"
	//	out, err := json.MarshalIndent(payload, "", "\t")
	//	if err != nil {
	//		app.errorLog.Println(err)
	//	}
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusOK)
	//	_, _ = w.Write(out)
	//})

	app.infoLog.Println("Listening on port ", app.cfg.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.cfg.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
