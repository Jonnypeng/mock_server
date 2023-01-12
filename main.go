package main

import (
	"io"
	"mock_server/controllers"
	"mock_server/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	// r.Use(middlewares.Cors())

	// 检验请求是否有合法token
	controllers.GetAuth(r)
	// 创建用户
	controllers.CreateUser(r)
	// 用户登录
	controllers.UserLogin(r)
	// 获取一名用户详情
	controllers.GetSingleUser(r)
	// 创建房屋
	controllers.CreateHouse(r)
	// 获取房屋详情
	controllers.GetSingleHouse(r)
	// 创建帖子
	controllers.CreatePost(r)
	// 获取一份帖
	controllers.GetSinglePost(r)
	// 获取一定数量的帖
	controllers.GetPostList(r)
	// 创建订单
	controllers.CreateOrder(r)
	// 获取订单详情
	controllers.GetSingleOrder(r)
	// 完成订单
	controllers.CompleteOrder(r)
	// 订单回调
	controllers.CashCallback(r)

	return r
}

func main() {
	cfg, err := utils.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	r := setupRouter()
	r.Run(":8008")
	r.Run(cfg.AppHost + ":" + cfg.AppPort)
}
