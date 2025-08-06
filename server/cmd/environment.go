// Package cmd The cmd package
package cmd

import (
	"os"
	"strconv"
)

// Variables The environment variables available for the project
type Variables struct {
	// Postgres Environment Varaibles
	POSTGRES_USER     string
	POSTGRES_DB       string
	POSTGRES_PASSWORD string
	POSTGRES_PORT     string

	// SERVER PORT
	SERVER_PORT string

	// ALGORITHM
	DEFAULT_REWARD float32
	DEFAULT_DELTA  int32

	LEARNING_RATE float32
	TEMPERATURE   float32
}

// SetupEnvironment Setup the environment varaibles that can be used throughout the application
func SetupEnvironment() *Variables {
	r, err := strconv.ParseFloat(os.Getenv("DEFAULT_REWARD"), 32)
	if err != nil {
		r = 0.5
	}

	d, err := strconv.ParseInt(os.Getenv("DEFAULT_DELTA"), 10, 32)
	if err != nil {
		d = 30
	}

	l, err := strconv.ParseFloat(os.Getenv("LEARNING_RATE"), 32)
	if err != nil {
		l = 0.1
	}

	t, err := strconv.ParseFloat(os.Getenv("TEMPERATURE"), 32)
	if err != nil {
		t = 0.5
	}

	env := &Variables{
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
		SERVER_PORT:       os.Getenv("SERVER_PORT"),
		DEFAULT_REWARD:    float32(r),
		DEFAULT_DELTA:     int32(d),
		LEARNING_RATE:     float32(l),
		TEMPERATURE:       float32(t),
	}

	return env
}
