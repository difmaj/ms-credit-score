package handler

import (
	"context"

	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/pkg/router"
	"github.com/difmaj/ms-credit-score/internal/pkg/router/middleware"
)

// IUsecase represents the interface for the usecase.
type IUsecase interface {
	// Login handles the login request.
	Login(context.Context, *dto.LoginHTTPInput) (*dto.LoginHTTPOutput, error)
}

// Handler represents the handler.
type Handler struct {
	usecase IUsecase
}

// NewHandler creates a new instance of Handler.
func NewHandler(router router.IRouter, middle *middleware.Middleware, usecase IUsecase) {
	handler := &Handler{usecase}

	{
		router.POST("login", handler.Login)
		// router.GET("health", handler.Health)
	}

	// context := "asset"
	// asset := router.Group(context, middle.BasicAuth())
	{ // asset
		// asset.GET("", middle.PermissionAuth(context, ""), handler.ListAssets)
		// asset.POST("", middle.PermissionAuth(context, "create"), handler.CreateAsset)
		// asset.GET(":asset_id", middle.PermissionAuth(context, "read"), handler.GetAssetByID)
		// asset.PATCH(":asset_id", middle.PermissionAuth(context, "update"), handler.UpdateAsset)
		// asset.DELETE(":asset_id", middle.PermissionAuth(context, "delete"), handler.DeleteAsset)
	}

	// context = "debt"
	// debt := router.Group(context, middle.BasicAuth())
	{ // debt
		// debt.GET("", middle.PermissionAuth(context, ""), handler.ListDebts)
		// debt.POST("", middle.PermissionAuth(context, "create"), handler.CreateDebt)
		// debt.GET(":debt_id", middle.PermissionAuth(context, "read"), handler.GetDebtByID)
		// debt.PATCH(":debt_id", middle.PermissionAuth(context, "update"), handler.UpdateDebt)
		// debt.DELETE(":debt_id", middle.PermissionAuth(context, "delete"), handler.DeleteDebt)
	}

	// context = "score"
	// score := router.Group(context, middle.BasicAuth())
	{ // score
		// score.GET(":user_id", middle.PermissionAuth(context, "read"), handler.GetScoreByUserID)
	}
}
