package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IRouter interface.
type IRouter interface {
	Engine() *gin.Engine
	Run(port int)
	Use(middleware ...gin.HandlerFunc)
	Group(path string, handlers ...gin.HandlerFunc) *gin.RouterGroup
	GET(path string, handlers ...gin.HandlerFunc)
	POST(path string, handlers ...gin.HandlerFunc)
	PUT(path string, handlers ...gin.HandlerFunc)
	PATCH(path string, handlers ...gin.HandlerFunc)
	DELETE(path string, handlers ...gin.HandlerFunc)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
