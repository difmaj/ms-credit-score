package migrations

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

// Config represents the configuration for the migrations.
type Config struct {
	// Path is the path of the migrations.
	Path string
	// Schema is the schema of the database.
	Schema string
}

// Result represents the result of the migrations.
type Result struct {
	// Version is the version of the migration.
	Version uint
	// Dirty is the dirty flag of the migration.
	Dirty bool
}

// Run runs the migrations.
func Run(orm *gorm.DB, config *Config) (*Result, error) {
	db, err := orm.DB()
	if err != nil {
		return nil, err
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: "schema_migrations",
	})
	if err != nil {
		return nil, err
	}

	migration, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", config.Path), "mysql", driver)
	if err != nil {
		return nil, err
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	version, dirty, err := migration.Version()
	if err != nil {
		return nil, err
	}
	return &Result{
		Version: version,
		Dirty:   dirty,
	}, nil
}
