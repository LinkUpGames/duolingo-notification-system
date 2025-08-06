// Package user The user information
package user

import (
	"fmt"
	"server/db"

	"github.com/google/uuid"
)

// GetUser Get the user given the id
func GetUser(db *db.DB, id string) map[string]any {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %s", id)

	user := db.GetEntry(query)

	_, ok := user["id"]

	if ok {
		return user
	} else {
		return nil
	}
}

func SetUser(db *db.DB, name string) (string, bool) {
	id := uuid.New().String()
	query := fmt.Sprintf("INSERT INTO users(id, name) WHERE (%s, %s);", id, name)

	status := db.SetEntry(query)

	return id, status
}
