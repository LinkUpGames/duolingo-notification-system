// Package user The user information
package user

import (
	"database/sql"
	"fmt"
	"server/db"

	"github.com/google/uuid"
)

// GetUser Get the user given the id
func GetUser(database *sql.DB, id string) map[string]any {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", id)

	user := db.GetEntry(database, query)

	_, ok := user["id"]

	if ok {
		return user
	} else {
		return nil
	}
}

func SetUser(database *sql.DB, name string) map[string]any {
	id := uuid.New().String()
	query := fmt.Sprintf("INSERT INTO users(id, name) WHERE (%s, %s);", id, name)

	// db.
}
