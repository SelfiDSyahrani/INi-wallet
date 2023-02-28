package handler

import (
	"dev_selfi/dto"
	"dev_selfi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	input := &dto.UserRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("register failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userUseCase.RegisterUser(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("register failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	token, err := h.jwtService.GenerateToken((newUser.UserWallet_ID))
	if err != nil {
		response := utils.ErrorResponse("register failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedLogin := dto.FormatLogin(newUser, token)
	response := utils.SuccessResponse("register success", http.StatusOK, formattedLogin)
	c.JSON(http.StatusOK, response)
}
