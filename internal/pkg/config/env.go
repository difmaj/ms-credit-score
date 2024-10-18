package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Environment is a struct that holds the environment variables.
type Environment struct {
	ServiceName string `env:"SERVICE_NAME" envDefault:"ms-task"`
	Port        string `env:"PORT" envDefault:"8080"`
	Environment string `env:"ENVIRONMENT" envDefault:"development"`

	DatabaseMYSQLDNS    string `env:"DATABASE_MYSQL_DNS"`
	DatabaseMYSQLSchema string `env:"DATABASE_MYSQL_SCHEMA"`
}

// Env is the environment variable.
var Env Environment

// Load loads the environment variables.
func Load(filenames ...string) error {
	err := godotenv.Load(filenames...)
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}

	Env = Environment{}
	if err := env.Parse(&Env); err != nil {
		return fmt.Errorf("failed to parse environment variables: %w", err)
	}
	return nil
}

func (env Environment) String() string {
	return env.Environment
}
