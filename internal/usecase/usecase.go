package usecase

import (
	"context"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

// IRepository interface.
type IRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetPrivilegesByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Privilege, error)
}

// IRedisClient represents the Redis client.
type IRedisClient interface {
	Client() redis.Conn
	Get(key string, v any) error
	Set(key string, v any, exp int) error
	ConnCheck() error
}

// Usecase struct.
type Usecase struct {
	repo  IRepository
	redis IRedisClient
}

// New creates a new Usecase.
func New(repo IRepository, redis IRedisClient) *Usecase {
	return &Usecase{repo: repo, redis: redis}
}
