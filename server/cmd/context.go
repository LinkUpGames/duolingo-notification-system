package cmd

import (
	"context"
	"server/db"
)

// Context keys
const (
	DATABASE = iota
	VARIABLES
)

// AppContext The context to be available throughout the whole application
type AppContext struct {
	// Context available through all application
	Ctx context.Context
}

// SetupContext Setup app context
func SetupContext() *AppContext {
	ctx := context.Background()

	// Setup the database
	db := db.SetupDatabase()
	ctx = context.WithValue(ctx, DATABASE, db)

	// Setup Variables
	v := SetupEnvironment()
	ctx = context.WithValue(ctx, VARIABLES, v)

	return &AppContext{
		Ctx: ctx,
	}
}
