package database

import (
	"github.com/difmaj/ms-credit-score/internal/pkg/config"
	"github.com/jmoiron/sqlx"
)

// Open opens a new database connection.
func Open() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", config.Env.DatabaseMYSQLDNS)
	if err != nil {
		return nil, err
	}
	return db, nil
}
