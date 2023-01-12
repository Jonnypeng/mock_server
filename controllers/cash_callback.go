package controllers

import (
	"fmt"
	"house_system_backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CashCallback(r *gin.Engine) {
	r.POST("cash/callback", func(c *gin.Context) {
		var callBackVo any
		c.ShouldBindJSON(&callBackVo)
		fmt.Println(callBackVo)
		c.JSON(http.StatusAccepted, model.ResponseData{
			Code: 200,
			Msg:  "",
		})
	})
}
