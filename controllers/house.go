package controllers

import (
	"fmt"
	"house_system_backend/model"
	"house_system_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHouse(r *gin.Engine) {
	r.POST("house/create", func(c *gin.Context) {
		var houseVo model.House
		jsonErr := c.ShouldBindJSON(&houseVo)
		if jsonErr != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  "参数错误",
			})
			return
		}
		// 创建房屋的sql语句
		sqlStr := "insert into House(license_number,user_id,city,region,title,imgs,house_type,storey,total_storey,decoration,isElevator,address,ownership_certificate,state) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
		res, err := utils.SqlDb.Exec(sqlStr, houseVo.LicenseNumber, houseVo.UserId, houseVo.City, houseVo.Region, houseVo.Title, houseVo.Imgs, houseVo.HouseType, houseVo.Storey, houseVo.TotalStorey, houseVo.Decoration, houseVo.IsElevator, houseVo.Address, houseVo.Ownership_certificate, houseVo.State)

		if err != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, model.ResponseData{
			Code: 200,
			Msg:  "房屋创建成功",
		})

		fmt.Println(res.RowsAffected())
	})
}

func GetSingleHouse(r *gin.Engine) {
	r.GET("house/info", func(c *gin.Context) {
		id := c.Query("id")
		sqlStr := "select address,city,Decoration,house_type ,id,imgs,isElevator,license_number,region,storey,title,total_storey,user_id,ownership_certificate,state from House where id = ?"
		var houseData model.House
		err := utils.SqlDb.QueryRow(sqlStr, id).Scan(
			&houseData.Address,
			&houseData.City,
			&houseData.Decoration,
			&houseData.HouseType,
			&houseData.Id,
			&houseData.Imgs,
			&houseData.IsElevator,
			&houseData.LicenseNumber,
			&houseData.Region,
			&houseData.Storey,
			&houseData.Title,
			&houseData.TotalStorey,
			&houseData.UserId,
			&houseData.Ownership_certificate,
			&houseData.State,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, model.ResponseData{
			Code: 200,
			Msg:  "获取成功",
			Data: houseData,
		})
	})
}
