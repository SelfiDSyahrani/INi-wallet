package route

import (
	"INi-Wallet/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

type RouterConfig struct {
}

func NewRouter(c *RouterConfig) *Router {
	return &Router{}
}

func (r *Router) Auth(route *gin.RouterGroup, h *handler.Handler) {
	route.POST("/sign-up", h.Register)
	route.POST("/sign-in", h.Login)
}
