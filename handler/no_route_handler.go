package handler

import (
	"INi-Wallet/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NoRoute(c *gin.Context) {
	response := utils.ErrorResponse("page not found", http.StatusNotFound, "page not found")
	c.JSON(http.StatusNotFound, response)
}
