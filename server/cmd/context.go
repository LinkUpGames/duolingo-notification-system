package cmd

import (
	"context"
	"server/db"
)

// Context keys
type key int

const (
	DATABASE key = iota
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

	// Setup Variables
	v := SetupEnvironment()
	ctx = context.WithValue(ctx, VARIABLES, v)

	// Setup the database
	db := db.Database(v.POSTGRES_DB, v.POSTGRES_USER, v.POSTGRES_PASSWORD, v.POSTGRES_PORT)
	ctx = context.WithValue(ctx, DATABASE, db)

	return &AppContext{
		Ctx: ctx,
	}
}
