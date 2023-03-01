package controller

import (
	"INi-Wallet/dto"
	"INi-Wallet/model"
	"INi-Wallet/service"
	"INi-Wallet/usecase"
	"INi-Wallet/utils"

	// "path/filepath"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUC usecase.UserUseCase
	jwt    service.JWTService
}

func (cc *UserController) loginUser(ctx *gin.Context) {
	input := &dto.LoginRequestBody{}
	err := ctx.ShouldBindJSON(input)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		user, err := cc.userUC.Login(input)
		if err != nil {
			utils.HandleInternalServerError(ctx, err.Error())
		}
		tokenUser, err := cc.jwt.GenerateToken(user.ID)
		if err != nil {
			utils.HandleInternalServerError(ctx, err.Error())
		} else {
			formatedLogin := dto.FormatLogin(&user, tokenUser)
			utils.HandleSuccess(ctx, "Login success", formatedLogin)
		}

	}

}

func (cc *UserController) registerCustomer(ctx *gin.Context) {
	var newuser model.User
	// file, errf := ctx.FormFile("file")
	// if errf != nil {
	// 	utils.HandleBadRequest(ctx, errf.Error())
	// 	return
	// }
	// filename := filepath.Base(file.Filename)
	// if errFileName := ctx.SaveUploadedFile(file, filename); errFileName != nil {
	// 	utils.HandleBadRequest(ctx, "Unable to upload file")
	// }
	err := ctx.ShouldBindJSON(&newuser)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		err := cc.userUC.RegisterUser(&newuser)
		if err != nil {
			utils.HandleInternalServerError(ctx, err.Error())
			// fmt.Println(newuser)
		} else {
			userTampilan := &dto.UserResponseBody{}
			userTampilan.Name = newuser.Name
			userTampilan.Email = newuser.Email
			userTampilan.ID = newuser.ID
			utils.HandleSuccessCreated(ctx, "Success Create New User", userTampilan)
		}
	}
}

func (cc *UserController) getAllUser(ctx *gin.Context) {
	users, err := cc.userUC.GetAllUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (cc *UserController) getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := cc.userUC.GetUserByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (cc *UserController) getByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := cc.userUC.GetByEmail(email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func NewUserController(router *gin.Engine, usecase usecase.UserUseCase) *UserController {
	newcontroller := UserController{
		userUC: usecase,
	}

	rG := router.Group("api/v1/INi-Wallet")
	rG.GET("/user", newcontroller.getAllUser)
	rG.GET("/user/:id", newcontroller.getUserById)
	rG.GET("/users/:email", newcontroller.getByEmail)
	rG.POST("/register", newcontroller.registerCustomer)
	rG.POST("/login", newcontroller.loginUser)
	return &newcontroller
}
