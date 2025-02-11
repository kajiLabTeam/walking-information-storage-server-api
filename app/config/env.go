package config

import (
	"os"

	_ "github.com/lib/pq"
)

type PostgresEnv struct {
	host     string
	port     string
	database string
	user     string
	password string
}

func NewPostgresEnv() *PostgresEnv {
	return &PostgresEnv{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		database: os.Getenv("DB_NAME"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
	}
}

func (pe *PostgresEnv) GetHostOfPrivateValue() string {
	return pe.host
}

func (pe *PostgresEnv) GetPortOfPrivateValue() string {
	return pe.port
}

func (pe *PostgresEnv) GetDatabaseOfPrivateValue() string {
	return pe.database
}

func (pe *PostgresEnv) GetUserOfPrivateValue() string {
	return pe.user
}

func (pe *PostgresEnv) GetPasswordOfPrivateValue() string {
	return pe.password
}
