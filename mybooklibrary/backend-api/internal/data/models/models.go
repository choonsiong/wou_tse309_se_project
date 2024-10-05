package models

import (
	"database/sql"
	"log"
	"time"
)

const databaseTimeout = time.Second * 3

var db *sql.DB

// Models type define database models used in the application.
type Models struct {
	Author    Author
	Book      Book
	Genre     Genre
	Publisher Publisher
	User      User
	UserBook  UserBook
	Token     Token
}

// New function returns initialized database models.
func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		Author:    Author{},
		Book:      Book{},
		Genre:     Genre{},
		Publisher: Publisher{},
		User:      User{},
		UserBook:  UserBook{},
		Token:     Token{},
	}
}

// CloseRows function ensures rows are closed properly or print an error message if not.
func CloseRows(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Printf("Error closing rows: %v", err)
	}
}
