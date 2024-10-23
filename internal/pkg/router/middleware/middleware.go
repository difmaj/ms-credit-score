package middleware

import "github.com/difmaj/ms-credit-score/internal/interfaces"

// Middleware struct.
type Middleware struct {
	uc interfaces.IUsecase
}

// NewMiddleware creates a new instance of the Middleware struct.
func NewMiddleware(uc interfaces.IUsecase) interfaces.IMiddleware {
	return &Middleware{uc}
}
