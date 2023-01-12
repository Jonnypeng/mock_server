package controllers

import (
	"fmt"
	"house_system_backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CashCallback(r *gin.Engine) {
	r.POST("cash/callback", func(c *gin.Context) {
		var callBackVo model.Callback
		if c.ShouldBindJSON(&callBackVo) != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  "参数错误",
			})
			return
		}
		fmt.Println(callBackVo)
		c.JSON(http.StatusBadRequest, model.ResponseData{
			Code: 200,
			Msg:  "",
		})
	})
}
