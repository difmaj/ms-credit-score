package repository

import (
	"database/sql"
)

// Repository represents the client for the MySQL database.
type Repository struct {
	db *sql.DB
}

// New creates a new instance of Client with the specified endpoint and secret.
func New(db *sql.DB) (*Repository, error) {
	return &Repository{db}, nil
}
