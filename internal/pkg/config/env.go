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

	JWTGlobalKey      string `env:"JWT_GLOBAL_KEY,required"`
	JWTAud            string `env:"JWT_AUD,required"`
	JWTIss            string `env:"JWT_ISS,required"`
	JWTRefreshDaysExp int    `env:"JWT_REFRESH_DAYS_EXP" envDefault:"2"`

	RedisMaxIdle     int `env:"REDIS_MAX_IDLE" envDefault:"50"`
	RedisMaxActive   int `env:"REDIS_MAX_ACTIVE" envDefault:"100"`
	RedisIdleTimeout int `env:"REDIS_IDLE_TIMEOUT" envDefault:"300"`

	RedisNetwork  string `env:"REDIS_NETWORK,required"`
	RedisAddr     string `env:"REDIS_ADDR,required"`
	RedisPassword string `env:"REDIS_PASSWORD,required"`
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
