// Package cmd The cmd package
package cmd

import (
	"os"
	"strconv"
)

// Variables The environment variables available for the project
type Variables struct {
	// Postgres Environment Varaibles
	PostgresUser     string
	PostgresDB       string
	PostgresPassword string
	PostgresPort     string

	// SERVER PORT
	ServerPort string

	// ALGORITHM
	DefaultReward bool
	CutOff        int
	Explore       float32
	Penalty       float32
	Factor        float32
	Score         float32
}

// SetupEnvironment Setup the environment varaibles that can be used throughout the application
func SetupEnvironment() *Variables {
	s, err := strconv.ParseFloat(os.Getenv("DEFAULT_SCORE"), 32)
	if err != nil {
		s = 0.5
	}

	c, err := strconv.ParseInt(os.Getenv("CUTOFF"), 10, 32)
	if err != nil {
		c = 10
	}

	p, err := strconv.ParseFloat(os.Getenv("PENALTY"), 32)
	if err != nil {
		p = 0.02
	}

	f, err := strconv.ParseFloat(os.Getenv("FACTOR"), 32)
	if err != nil {
		f = 10
	}

	e, err := strconv.ParseFloat(os.Getenv("EXPLORE"), 32)
	if err != nil {
		e = 1.2
	}

	env := &Variables{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		ServerPort:       os.Getenv("SERVER_PORT"),
		Score:            float32(s),
		CutOff:           int(c),
		Penalty:          float32(p),
		Factor:           float32(f),
		Explore:          float32(e),
	}

	return env
}
