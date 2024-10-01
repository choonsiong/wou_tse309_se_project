package main

import (
	"app-backend/internal/driver"
	"database/sql"
	"log"
	"os"
)

var (
	cfg Config
)

// main function is the entry point of the application
func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	err := loadConfig(&cfg)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := driver.ConnectPostgresDatabase(cfg.DSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {
			log.Fatalf("Failed to close SQL connection: %v", err)
		}
	}(db.SQL)

	app := &application{
		cfg:      cfg,
		db:       db,
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	log.Fatal(app.serve())
}
