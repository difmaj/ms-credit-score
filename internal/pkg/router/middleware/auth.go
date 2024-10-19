package middleware

import (
	"net/http"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/pkg/router/response"
	"github.com/gin-gonic/gin"
)

// BasicAuth is a middleware that checks if the user is authenticated.
func (m *Middleware) BasicAuth() func(*gin.Context) {
	return func(ctx *gin.Context) {
		user, err := m.checkToken(ctx)
		if err != nil {
			response.Error(ctx.Writer, dto.NewAPIError(http.StatusUnauthorized, err, ""))
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}

// PermissionAuth is a middleware that checks if the user has the permission to access the endpoint.
func (m *Middleware) PermissionAuth(context string, action string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*domain.Claims)
		if err := m.checkPermissions(user, context, action); err != nil {
			response.Error(ctx.Writer, &dto.APIError{
				Status:  http.StatusUnauthorized,
				Message: "permission denied",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func (m *Middleware) checkPermissions(claims *domain.Claims, context, action string) error {
	hasPermission := false
	actions, ok := claims.Permissions[context]
	if ok {
		for _, act := range actions {
			if act == action || action == "" {
				hasPermission = true
				break
			}
		}
	}

	if !hasPermission {
		return &dto.APIError{
			Status:  http.StatusUnauthorized,
			Message: "permission denied",
		}
	}
	return nil
}

func (m *Middleware) checkToken(ctx *gin.Context) (*domain.Claims, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return nil, &dto.APIError{
			Status:  http.StatusUnauthorized,
			Message: "missing token",
		}
	}

	claims, err := m.uc.ClaimsJWT(token)
	if err != nil {
		return nil, err
	}
	if err := claims.Valid(); err != nil {
		return nil, err
	}
	return claims, nil
}
