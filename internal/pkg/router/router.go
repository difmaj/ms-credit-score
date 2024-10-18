package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/difmaj/ms-credit-score/internal/pkg/config"
	"github.com/difmaj/ms-credit-score/internal/pkg/logger"
	"github.com/difmaj/ms-credit-score/internal/pkg/validator"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router struct.
type Router struct {
	engine *gin.Engine
}

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

// NewRouter creates a new instance of the Router struct.
func NewRouter(middlewares ...gin.HandlerFunc) IRouter {

	engine := gin.New()
	engine.RedirectFixedPath = false
	engine.RedirectTrailingSlash = false
	engine.RemoveExtraSlash = true

	switch config.Env.Environment {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	binding.Validator = new(validator.DefaultValidator)

	if middlewares != nil {
		engine.Use(middlewares...)
	}

	engine.Use(ginzap.RecoveryWithZap(logger.Logger, true))
	engine.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, true))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return Router{engine: engine}
}

// Engine returns the gin.Engine.
func (r Router) Engine() *gin.Engine {
	return r.engine
}

// Run starts the server.
func (r Router) Run(port int) {
	fmt.Printf("Running on Port :%d", port)
	err := r.engine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err.Error())
	}
}

// Use adds middleware to the router.
func (r Router) Use(middleware ...gin.HandlerFunc) {
	r.engine.Use(middleware...)
}

// Group creates a new group of routes.
func (r Router) Group(path string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return r.engine.Group(path, handlers...)
}

// GET adds a GET route to the router.
func (r Router) GET(path string, handlers ...gin.HandlerFunc) {
	r.engine.GET(path, handlers...)
}

// POST adds a POST route to the router.
func (r Router) POST(path string, handlers ...gin.HandlerFunc) {
	r.engine.POST(path, handlers...)
}

// PUT adds a PUT route to the router.
func (r Router) PUT(path string, handlers ...gin.HandlerFunc) {
	r.engine.PUT(path, handlers...)
}

// PATCH adds a PATCH route to the router.
func (r Router) PATCH(path string, handlers ...gin.HandlerFunc) {
	r.engine.PATCH(path, handlers...)
}

// DELETE adds a DELETE route to the router.
func (r Router) DELETE(path string, handlers ...gin.HandlerFunc) {
	r.engine.DELETE(path, handlers...)
}

// ServeHTTP serves the HTTP request.
func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.engine.ServeHTTP(w, req)
}
