package main

import (
	"log"
	"os"
)

var (
	cfg config
)

// main function is the entry point of the application
func main() {
	cfg.port = 8081

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := &application{
		cfg:      cfg,
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	log.Fatal(app.serve())
}
