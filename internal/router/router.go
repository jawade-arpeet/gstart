package router

import (
	"gstart/internal/constants"
	"gstart/internal/handler"
	"gstart/internal/middleware"
	v1 "gstart/internal/router/v1"

	"github.com/gin-gonic/gin"
)

func New(
	runEnv constants.Env,
	middleware *middleware.Middleware,
	handler *handler.Handler,
) *gin.Engine {
	if runEnv == constants.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	apiGrp := router.Group("/api")

	v1.MountV1Router(apiGrp, middleware, handler)

	return router
}
