package middleware

import (
	"github.com/difmaj/ms-credit-score/internal/domain"
)

// IUsecase interface.
type IUsecase interface {
	ClaimsJWT(token string) (*domain.Claims, error)
}

// Middleware struct.
type Middleware struct {
	uc IUsecase
}

// NewMiddleware creates a new instance of the Middleware struct.
func NewMiddleware(uc IUsecase) *Middleware {
	return &Middleware{uc}
}
