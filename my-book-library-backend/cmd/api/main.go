package main

import (
	"app-backend/internal/driver"
	"log"
	"os"
)

var (
	cfg config
)

// main function is the entry point of the application
func main() {
	cfg.port = 9009

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	dsn := os.Getenv("DSN")

	db, err := driver.ConnectPostgresDatabase(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := &application{
		cfg:      cfg,
		db:       db,
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	log.Fatal(app.serve())
}
