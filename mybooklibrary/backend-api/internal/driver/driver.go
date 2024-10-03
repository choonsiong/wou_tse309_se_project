package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

type DB struct {
	SQL *sql.DB
}

var dbConnection = &DB{}

const maxDatabaseLifeTime = time.Minute * 5
const maxIdleDatabaseConnection = 5
const maxOpenDatabaseConnection = 5

func ConnectPostgresDatabase(dsn string) (*DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(maxDatabaseLifeTime)
	db.SetMaxIdleConns(maxIdleDatabaseConnection)
	db.SetMaxOpenConns(maxOpenDatabaseConnection)

	err = testDatabaseConnection(db)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	dbConnection.SQL = db

	return dbConnection, nil
}

func testDatabaseConnection(db *sql.DB) error {
	return db.Ping()
}
