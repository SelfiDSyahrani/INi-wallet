package handler

import (
	"INi-Wallet/dto"
	"INi-Wallet/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	input := &dto.RegisterRequestBody{}

	err := c.ShouldBindJSON(input)
	log.Println(err)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("register failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// err := h.userUseCase.RegisterUser(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("register failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}
	// token, err := h.jwtService.GenerateToken(newUser.UserWallet_ID)
	if err != nil {
		response := utils.ErrorResponse("register failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// formattedLogin := dto.FormatLogin(newUser, token)
	// response := utils.SuccessResponse("register success", http.StatusOK, formattedLogin)
	// c.JSON(http.StatusOK, response)
}

func (h *Handler) Login(c *gin.Context) {
	input := &dto.LoginRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("login failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userUseCase.Attempt(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("login failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	token, err := h.jwtService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := utils.ErrorResponse("login failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedLogin := dto.FormatLogin(loggedinUser, token)
	response := utils.SuccessResponse("login success", http.StatusOK, formattedLogin)
	c.JSON(http.StatusOK, response)
}
