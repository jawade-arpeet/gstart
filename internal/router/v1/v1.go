package v1

import (
	"gstart/internal/handler"
	"gstart/internal/middleware"

	"github.com/gin-gonic/gin"
)

func MountV1Router(
	routerGrp *gin.RouterGroup,
	middleware *middleware.Middleware,
	handler *handler.Handler,
) {
	v1Grp := routerGrp.Group("/v1")

	mountHealthRouter(v1Grp, handler.Health)
}
