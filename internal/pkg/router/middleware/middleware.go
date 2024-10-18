package middleware

// IUsecase interface.
type IUsecase interface{}

// Middleware struct.
type Middleware struct {
	uc IUsecase
}

// NewMiddleware creates a new instance of the Middleware struct.
func NewMiddleware(uc IUsecase) *Middleware {
	return &Middleware{uc}
}
