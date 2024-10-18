package redis

import (
	"time"

	"github.com/difmaj/ms-credit-score/internal/pkg/config"
	"github.com/gomodule/redigo/redis"
)

// Client represents the Redis client.
type Client struct {
	pool *redis.Pool
}

// NewClient creates a new instance of Client.
func NewClient() (*Client, error) {
	pool := &redis.Pool{
		MaxIdle:     config.Env.RedisMaxIdle,
		MaxActive:   config.Env.RedisMaxActive,
		IdleTimeout: time.Duration(config.Env.RedisIdleTimeout) * time.Millisecond,
		Dial: func() (redis.Conn, error) {
			client, err := redis.Dial(config.Env.RedisNetwork, config.Env.RedisAddr)
			if err != nil {
				return nil, err
			}

			if _, err := client.Do("AUTH", config.Env.RedisPassword); err != nil {
				client.Close()
				return nil, err
			}

			return client, err
		},
	}

	return &Client{
		pool: pool,
	}, nil
}

// Client returns a connection to the Redis server.
func (rc *Client) Client() redis.Conn {
	return rc.pool.Get()
}

// Get gets a value from the Redis server.
func (rc *Client) Get(key string, v any) error {
	conn := rc.Client()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return err
	}

	*v.(*string) = value
	return nil
}

// Set sets a value in the Redis server.
func (rc *Client) Set(key string, v any, exp int) error {
	conn := rc.Client()
	defer conn.Close()

	_, err := conn.Do("SET", key, v)
	if err != nil {
		return err
	}

	if exp > 0 {
		_, err = conn.Do("EXPIRE", key, exp)
	}
	return err
}

// ConnCheck checks the connection to the Redis server.
func (rc *Client) ConnCheck() error {
	conn := rc.Client()
	defer conn.Close()

	_, err := conn.Do("PING")
	return err
}
