package controller

import (
	"INi-Wallet/model"
	"INi-Wallet/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router  *gin.Engine
	usecase usecase.UserUseCase
}

func (cc *UserController) registerUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := cc.usecase.RegisterUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (cc *UserController) getAllUser(ctx *gin.Context) {
	users, err := cc.usecase.GetAllUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (cc *UserController) getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := cc.usecase.GetUserByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func NewCustomerController(r *gin.Engine, usecase usecase.UserUseCase) *UserController {
	controller := UserController{
		router:  r,
		usecase: usecase,
	}
	r.GET("/Profile", controller.getAllUser)
	r.GET("/Profile/:id", controller.getUserById)
	r.POST("/Register", controller.registerUser)
	return &controller
}
