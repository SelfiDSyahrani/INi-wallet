package controller

import (
	"dev_selfi/usecase"

	"github.com/gin-gonic/gin"
)

type MoneyChangerController struct {
	router  *gin.Engine
	usecase usecase.MoneyChangerUsecase
}

func (mcc * MoneyChangerController) GetAllMoneyChangerController(c *gin.Context) {
	MoneyChanger,err := mcc.usecase.MoneyChangerAll()
	if err != nil {
		c.JSON(400,gin.H{
			"message" : err.Error(),
		})

		return
	}
	
		c.JSON(200,gin.H{
			"message" : "OK",
			"data" : MoneyChanger,
		})
}

func NewMoneyChangerController(mcc *gin.Engine,usecase usecase.MoneyChangerUsecase)  {
	controller :=MoneyChangerController{
		router: mcc,
		usecase: usecase,
	}
	mcc.GET("/MoneyChanger",controller.GetAllMoneyChangerController)
}