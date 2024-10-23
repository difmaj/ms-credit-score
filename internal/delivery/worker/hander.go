package handler

import (
	"context"

	"github.com/difmaj/ms-credit-score/internal/interfaces"
)

// Handler represents the handler.
type Handler struct {
	usecase interfaces.IUsecase
}

// NewHandler creates a new instance of Handler.
func NewHandler(subscriber interfaces.ISubscriber, middle interfaces.IMiddleware, usecase interfaces.IUsecase) {
	ctx := context.Background()
	handler := &Handler{usecase}
	{
		subscriber.Subscribe(ctx, "login.topic", handler.Login)
	}

	context := "asset"
	{ // asset
		subscriber.Subscribe(ctx, "", middle.BasicAuth(), middle.PermissionAuth(context, "list"), handler.ListAssets)
		subscriber.Subscribe(ctx, "", middle.BasicAuth(), middle.PermissionAuth(context, "create"), handler.CreateAsset)
		subscriber.Subscribe(ctx, ":asset_id", middle.BasicAuth(), middle.PermissionAuth(context, "read"), handler.GetAssetByID)
		subscriber.Subscribe(ctx, ":asset_id", middle.BasicAuth(), middle.PermissionAuth(context, "update"), handler.UpdateAsset)
		subscriber.Subscribe(ctx, ":asset_id", middle.BasicAuth(), middle.PermissionAuth(context, "delete"), handler.DeleteAsset)
	}

	context = "debt"
	{ // debt
		subscriber.Subscribe(ctx, "", middle.BasicAuth(), middle.PermissionAuth(context, "list"), handler.ListDebts)
		subscriber.Subscribe(ctx, "", middle.BasicAuth(), middle.PermissionAuth(context, "create"), handler.CreateDebt)
		subscriber.Subscribe(ctx, ":debt_id", middle.BasicAuth(), middle.PermissionAuth(context, "read"), handler.GetDebtByID)
		subscriber.Subscribe(ctx, ":debt_id", middle.BasicAuth(), middle.PermissionAuth(context, "update"), handler.UpdateDebt)
		subscriber.Subscribe(ctx, ":debt_id", middle.BasicAuth(), middle.PermissionAuth(context, "delete"), handler.DeleteDebt)
	}

	// context = "score"
	// score := router.Group(context, middle.BasicAuth())
	// { // score
	// 	score.GET(":user_id", middle.PermissionAuth(context, "read"), handler.GetScoreByUserID)
	// }
}
