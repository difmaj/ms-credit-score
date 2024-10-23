package handler

import (
	"net/http"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/pkg/router/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAssetByID returns a asset by asset ID.
func (h *Handler) GetAssetByID(ctx *gin.Context) {
	user := ctx.MustGet("user").(*domain.Claims)

	request := new(dto.GetAssetByIDInput)
	if err := ctx.ShouldBindUri(request); err != nil {
		ctx.Error(err)
		return
	}

	resp, err := h.usecase.GetAssetByID(ctx, user.User.ID, request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Ok(ctx.Writer, http.StatusOK, resp)
}

// ListAssets returns the assets of a user.
func (h *Handler) ListAssets(ctx *gin.Context) {
	user := ctx.MustGet("user").(*domain.Claims)

	resp, err := h.usecase.GetAssetsByUserID(ctx, user.User.ID)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Ok(ctx.Writer, http.StatusOK, resp)
}

// CreateAsset creates a new asset.
func (h *Handler) CreateAsset(ctx *gin.Context) {
	user := ctx.MustGet("user").(*domain.Claims)

	request := new(dto.CreateAssetInput)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.Error(err)
		return
	}

	resp, err := h.usecase.CreateAsset(ctx, user.User.ID, request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Ok(ctx.Writer, http.StatusCreated, resp)
}

// UpdateAsset updates an asset.
func (h *Handler) UpdateAsset(ctx *gin.Context) {
	user := ctx.MustGet("user").(*domain.Claims)

	assetID, err := uuid.Parse(ctx.Param("asset_id"))
	if err != nil {
		ctx.Error(err)
		return
	}

	request := new(dto.UpdateAssetInput)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.Error(err)
		return
	}

	resp, err := h.usecase.UpdateAsset(ctx, user.User.ID, assetID, request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Ok(ctx.Writer, http.StatusOK, resp)
}

// DeleteAsset deletes an asset.
func (h *Handler) DeleteAsset(ctx *gin.Context) {
	user := ctx.MustGet("user").(*domain.Claims)

	request := new(dto.DeleteAssetInput)
	if err := ctx.ShouldBindUri(request); err != nil {
		ctx.Error(err)
		return
	}

	err := h.usecase.DeleteAsset(ctx, user.User.ID, request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Ok(ctx.Writer, http.StatusNoContent, response.Empty{})
}
