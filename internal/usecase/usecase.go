package usecase

import "github.com/difmaj/ms-credit-score/internal/interfaces"

// Usecase struct.
type Usecase struct {
	repo  interfaces.IRepository
	redis interfaces.IRedisClient
}

// New creates a new Usecase.
func New(repo interfaces.IRepository, redis interfaces.IRedisClient) interfaces.IUsecase {
	return &Usecase{repo: repo, redis: redis}
}
