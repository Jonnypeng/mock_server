package controllers

import (
	"fmt"
	"house_system_backend/model"
	"house_system_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建用户模块，传入gin的引擎实例
func CreateUser(r *gin.Engine) {
	// 路由控制注册字段
	r.POST("user/register", func(c *gin.Context) {
		var userVo model.User
		// 参数拦截
		if c.ShouldBindJSON(&userVo) != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  "参数错误",
			})
			return
		}
		// sql语句进行插入
		sqlStr := "insert into Users(user_name, user_img, phone, password) values(?,?,?,?)"
		res, err := utils.SqlDb.Exec(sqlStr, userVo.UserName, userVo.UserImg, userVo.Phone, userVo.Password)
		// 数据库操作拦截
		if err != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  err.Error(),
			})
			return
		}

		// 成功创建数据的响应操作
		c.JSON(http.StatusOK, model.ResponseData{
			Code: 200,
			Msg:  "注册成功",
		})

		fmt.Println(res.RowsAffected())

	})
}

// 用户登录
func UserLogin(r *gin.Engine) {
	r.POST("user/login", func(c *gin.Context) {
		var userVo model.User
		if c.ShouldBindJSON(&userVo) != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  "参数错误",
			})
			return
		}
		sqlStr := "select id, user_name,phone,password from Users where phone = ?"
		var dbUser = model.User{}
		// 使用用户的电话号码进行数据库查询
		err := utils.SqlDb.QueryRow(sqlStr, userVo.Phone).Scan(&dbUser.Id, &dbUser.UserName, &dbUser.Phone, &dbUser.Password)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  err.Error(),
			})
			return
		}
		// 如果数据库没有此用户，或此用户的电话号码与密码与数据库中存储的不一致，将会返回登录失败
		if userVo.Phone == dbUser.Phone && userVo.Password == dbUser.Password {
			info := model.NewUserInfo(dbUser)
			tokenString, _ := GenerateToken(*info)
			c.JSON(http.StatusOK, model.ResponseData{
				Code: 200,
				Msg:  "登录成功",
				Data: map[string]any{
					"token":   tokenString,
					"user_id": dbUser.Id,
				},
			})
			return
		}
		c.String(http.StatusUnauthorized, "登录失败")
	})
}

// 获取一个用户
func GetSingleUser(r *gin.Engine) {
	r.GET("user/info", func(c *gin.Context) {
		id := c.Query("id")
		sqlStr := "select id,user_img,phone,user_name from Users where id = ?"
		var userData model.UserInfo
		err := utils.SqlDb.QueryRow(sqlStr, id).Scan(&userData.Id, &userData.UserImg, &userData.Phone, &userData.UserName)
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
			Data: userData,
		})
	})
}
