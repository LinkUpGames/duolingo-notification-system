// Package user The user information
package user

import (
	"fmt"
	"server/db"

	"github.com/google/uuid"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateUser Create a new user struct and returns a pointer
func CreateUser(id string, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}

// GetUser Get the user given the id
func GetUser(db *db.DB, id string) *User {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = '%s'", id)

	user := db.GetEntry(query)

	_, ok := user["id"]

	if ok {
		name := user["name"].(string)

		return CreateUser(id, name)
	} else {
		return nil
	}
}

// SetUser Set a new user on the database. Returns the id of the user created if any
func SetUser(db *db.DB, name string) string {
	id := uuid.New().String()
	query := fmt.Sprintf("INSERT INTO users(id, name) WHERE ('%s', '%s');", id, name)

	status := db.SetEntry(query)

	if !status {
		id = ""
	}

	return id
}
