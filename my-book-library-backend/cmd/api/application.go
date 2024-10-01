package main

import (
	"app-backend/internal/data/models"
	"fmt"
	"log"
	"net/http"
)

// application type store application wide shared data
type application struct {
	cfg      Config
	models   models.Models
	errorLog *log.Logger
	infoLog  *log.Logger
}

// start the web server
func (app *application) serve() error {
	app.infoLog.Println("Listening on port ", app.cfg.HttpPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.cfg.HttpPort),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
