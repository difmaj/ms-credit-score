package handler

import (
	"context"

	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/pkg/router"
	"github.com/difmaj/ms-credit-score/internal/pkg/router/middleware"
)

// IUsecase represents the interface for the v6dataviz usecase.
type IUsecase interface {
	// Login handles the login request.
	Login(context.Context, *dto.LoginHTTPInput) (*dto.LoginHTTPOutput, error)
}

// Handler represents the handler for the v6dataviz group.
type Handler struct {
	usecase IUsecase
}

// NewHandler creates a new instance of HandlerV6Dataviz.
func NewHandler(router router.IRouter, middle *middleware.Middleware, usecase IUsecase) {
	handler := &Handler{usecase}
	router.POST("login", handler.Login)
}
