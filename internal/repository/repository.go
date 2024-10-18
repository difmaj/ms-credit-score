package repository

import "github.com/jmoiron/sqlx"

// Repository represents the client for the MySQL database.
type Repository struct {
	db *sqlx.DB
}

// New creates a new instance of Client with the specified endpoint and secret.
func New(db *sqlx.DB) (*Repository, error) {
	return &Repository{db}, nil
}
