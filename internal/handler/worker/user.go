package handler

import (
	"net/http"

	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/pkg/router/response"
	"github.com/gin-gonic/gin"
)

// Login handles the login request.
func (h *Handler) Login(ctx *gin.Context) {
	request := new(dto.LoginHTTPInput)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.Error(err)
		return
	}

	resp, err := h.usecase.Login(ctx, request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Ok(ctx.Writer, http.StatusOK, resp)
}
