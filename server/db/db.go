// Package db is the package for connecting with the database and provides the operation interface for fetching and storing information
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB The database
type DB struct {
	client *sql.DB
}

// SetupDatabase Setup the connection with the postgres sql database and return it
func setupDatabase(dbname string, user string, password string, port string) *sql.DB {
	connection := fmt.Sprintf("user=%s dbname=%s host=database password=%s port=%s sslmode=disable", user, dbname, password, port)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		message := fmt.Sprintf("Error connecting to database: %s", err)
		panic(message)
	}

	return db
}

// Database Setup and return a working database
func Database(dbname string, user string, password string, port string) *DB {
	postgres := setupDatabase(dbname, user, password, port)

	db := &DB{
		client: postgres,
	}

	return db
}
