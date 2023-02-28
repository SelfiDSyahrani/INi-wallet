package route

import (
	"dev_selfi/handler"

	"github.com/gin-gonic/gin"
)

func (r *Router) Auth(route *gin.RouterGroup, h *handler.Handler) {
	route.POST("/sign-up", h.Register)
}
