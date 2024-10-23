package interfaces

import "github.com/gomodule/redigo/redis"

// IRedisClient represents the Redis client.
type IRedisClient interface {
	Client() redis.Conn
	Get(key string, v any) error
	Set(key string, v any, exp int) error
	ConnCheck() error
}
