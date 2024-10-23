package interfaces

import (
	"github.com/gin-gonic/gin"
)

// IMiddleware represents the middleware interface.
type IMiddleware interface {
	// ErrorHandler returns a middleware that handles errors.
	ErrorHandler() gin.HandlerFunc
	// BasicAuth returns a middleware that handles basic authentication.
	BasicAuth() func(*gin.Context)
	// PermissionAuth returns a middleware that handles permission-based authentication.
	PermissionAuth(context string, action string) func(*gin.Context)
}
