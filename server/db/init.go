// db is the package for connecting with the database and provides the operation interface for fetching and storing information
package db

import (
	"database/sql"
	"fmt"
	"os"
)

// SetupDatabase Setup the connection with the postgres sql database and return it
func SetupDatabase() *sql.DB {
	dbname := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")

	connection := fmt.Sprintf("user=%s dbname=%s host=database password=%s port=%s sslmode=disable", user, dbname, password, port)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		message := fmt.Sprintf("Error connecting to database: %s", err)
		panic(message)
	}

	return db
}
