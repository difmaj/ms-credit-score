package middleware

import (
	"errors"
	"net/http"

	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/pkg/router/response"
	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware that handles the errors.
func (m *Middleware) ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		var status int

		errs := make([]*dto.APIError, 0, len(ctx.Errors))
		for _, err := range ctx.Errors {
			var exs dto.APIErrors
			apiError := &dto.APIError{}
			if errors.As(err.Err, &apiError) {
				if status == 0 {
					status = apiError.Status
				}
				errs = append(errs, apiError)
			} else if errors.As(err.Err, &exs) {
				if status == 0 {
					status = http.StatusBadRequest
				}
				errs = append(errs, exs...)
			} else {
				errs = append(errs, &dto.APIError{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
					Err:     err,
				})
			}
		}

		if status == 0 {
			status = http.StatusInternalServerError
		}

		response.Error(ctx.Writer, errs...)
		ctx.Abort()
	}
}
