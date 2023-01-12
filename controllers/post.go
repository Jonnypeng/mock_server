package controllers

import (
	"fmt"
	"house_system_backend/model"
	"house_system_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(r *gin.Engine) {
	r.POST("post/create", func(c *gin.Context) {
		var postVo model.Post
		if c.ShouldBindJSON(&postVo) != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  "参数错误",
			})
			return
		}
		sqlStr := "insert into Post(license_number,user_id,post_type,create_time,price,state) values(?,?,?,?,?,?)"
		res, err := utils.SqlDb.Exec(sqlStr, postVo.LicenseNumber, postVo.UserId, postVo.PostType, postVo.CreateTime, postVo.Price, postVo.State)
		if err != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, model.ResponseData{
			Code: 200,
			Msg:  "创建成功",
		})
		fmt.Println(res.RowsAffected())
	})
}

// 获取一份帖
func GetSinglePost(r *gin.Engine) {
	r.GET("post/info", func(c *gin.Context) {
		id := c.Query("id")
		// sqlStr := "select id, license_number,user_id, post_type, create_time, price, state from Post where id = ?"
		sqlStr := "SELECT SP.*,House.imgs,House.title,House.house_type,House.id as house_id from (SELECT * from Post where id = ?) as SP LEFT JOIN House ON SP.license_number = House.license_number;"
		var postData model.Post
		err := utils.SqlDb.QueryRow(sqlStr, id).Scan(&postData.Id, &postData.LicenseNumber, &postData.UserId, &postData.PostType, &postData.CreateTime, &postData.Price, &postData.State, &postData.Imgs, &postData.Title, &postData.HouseType, &postData.HouseId)
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
			Data: postData,
		})
	})
}

func GetPostList(r *gin.Engine) {
	r.GET("post/list", func(c *gin.Context) {
		// 最大多少数量的数据
		limit := c.Query("limit")
		// 游标位置
		cursor := c.Query("cursor")
		// sqlStr := "select * from Post order by id desc limit ?,? "
		sqlStr := "SELECT Post.*,House.imgs,House.title,House.house_type,House.id as house_id from Post LEFT JOIN House ON Post.license_number = House.license_number order by id desc limit ? , ?"
		rows, err := utils.SqlDb.Query(sqlStr, cursor, limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  err.Error(),
			})
			return
		}
		defer rows.Close()
		// 做一个临时切片存储数据
		listPost := make([]model.Post, 0)
		// 迭代器向listPost注入数据
		for rows.Next() {
			var postData model.Post
			rows.Scan(&postData.Id, &postData.LicenseNumber, &postData.UserId, &postData.PostType, &postData.CreateTime, &postData.Price, &postData.State, &postData.Imgs, &postData.Title, &postData.HouseType, &postData.HouseId)
			listPost = append(listPost, postData)
		}
		c.JSON(http.StatusOK, model.ResponseData{
			Code: 200,
			Msg:  "获取成功",
			Data: listPost,
		})
	})
}
