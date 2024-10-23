package handler

import (
	"net/http"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/pkg/router/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetDebtByID returns a debt by debt ID.
func (h *Handler) GetDebtByID(ctx *gin.Context) {
	user := ctx.MustGet("user").(*domain.Claims)

	request := new(dto.GetDebtByIDInput)
	if err := ctx.ShouldBindUri(request); err != nil {
		ctx.Error(err)
		return
	}

	resp, err := h.usecase.GetDebtByID(ctx, user.User.ID, request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Ok(ctx.Writer, http.StatusOK, resp)
}

// ListDebts returns the debts of a user.
func (h *Handler) ListDebts(ctx *gin.Context) {
	user := ctx.MustGet("user").(*domain.Claims)

	resp, err := h.usecase.GetDebtsByUserID(ctx, user.User.ID)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Ok(ctx.Writer, http.StatusOK, resp)
}

// CreateDebt creates a new debt.
func (h *Handler) CreateDebt(ctx *gin.Context) {
	user := ctx.MustGet("user").(*domain.Claims)

	request := new(dto.CreateDebtInput)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.Error(err)
		return
	}

	resp, err := h.usecase.CreateDebt(ctx, user.User.ID, request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Ok(ctx.Writer, http.StatusCreated, resp)
}

// UpdateDebt updates an debt.
func (h *Handler) UpdateDebt(ctx *gin.Context) {
	user := ctx.MustGet("user").(*domain.Claims)

	DebtID, err := uuid.Parse(ctx.Param("Debt_id"))
	if err != nil {
		ctx.Error(err)
		return
	}

	request := new(dto.UpdateDebtInput)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.Error(err)
		return
	}

	resp, err := h.usecase.UpdateDebt(ctx, user.User.ID, DebtID, request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Ok(ctx.Writer, http.StatusOK, resp)
}

// DeleteDebt deletes an debt.
func (h *Handler) DeleteDebt(ctx *gin.Context) {
	user := ctx.MustGet("user").(*domain.Claims)

	request := new(dto.DeleteDebtInput)
	if err := ctx.ShouldBindUri(request); err != nil {
		ctx.Error(err)
		return
	}

	err := h.usecase.DeleteDebt(ctx, user.User.ID, request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Ok(ctx.Writer, http.StatusNoContent, response.Empty{})
}
