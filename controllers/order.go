package controllers

import (
	"fmt"
	"mock_server/model"
	"mock_server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 新建订单
func CreateOrder(r *gin.Engine) {
	r.POST("order/create", func(c *gin.Context) {
		var orderVo model.Order
		if c.ShouldBindJSON(&orderVo) != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  "参数错误",
			})
			return
		}
		sqlStr := "insert into Orders(post_id,house_id,user_id,buy_user_id,state,price) values(?,?,?,?,?,?)"
		res, err := utils.SqlDb.Exec(sqlStr, orderVo.PostId, orderVo.HouseId, orderVo.UserId, orderVo.BuyUserId, orderVo.State, orderVo.Price)
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

// 完成订单
func CompleteOrder(r *gin.Engine) {
	r.POST("order/complete", func(c *gin.Context) {
		var reqData model.ReqCompleteOrder
		if c.ShouldBindJSON(&reqData) != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  "参数错误",
			})
			return
		}
		// 更新订单数据
		sqlOrderStr := "update Orders SET state = 2 WHERE id = ?;"
		resOrder, errOrder := utils.SqlDb.Exec(sqlOrderStr, reqData.OrderId)

		// 更新帖数据
		sqlPostStr := "update Post SET state = 2 WHERE id = ?;"
		resPost, errPost := utils.SqlDb.Exec(sqlPostStr, reqData.PostId)

		// 更新房屋数据
		sqlHouseStr := "update House SET state = 2 WHERE id = ?;"
		resHouse, errHouse := utils.SqlDb.Exec(sqlHouseStr, reqData.HouseId)

		if errOrder != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  errOrder.Error(),
			})
			return
		}

		if errPost != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  errPost.Error(),
			})
			return
		}

		if errHouse != nil {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  errHouse.Error(),
			})
			return
		}

		c1, _ := resOrder.RowsAffected()
		c2, _ := resOrder.RowsAffected()
		c3, _ := resOrder.RowsAffected()

		count := c1 + c2 + c3

		// 只有三条数据都更新才能判断，这是一个成功的动作
		if count != 3 {
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 400,
				Msg:  "更新错误",
			})
			return
		}

		c.JSON(http.StatusOK, model.ResponseData{
			Code: 200,
			Msg:  "更新成功",
		})

		fmt.Println(resOrder.RowsAffected())
		fmt.Println(resPost.RowsAffected())
		fmt.Println(resHouse.RowsAffected())
	})
}

// 获取订单详情
func GetSingleOrder(r *gin.Engine) {
	r.GET("order/info", func(c *gin.Context) {
		id := c.Query("id")
		sqlStr := "select id, post_id, house_id, user_id, buy_user_id, state, price from Orders where id = ?"
		var orderData model.Order
		err := utils.SqlDb.QueryRow(sqlStr, id).Scan(&orderData.Id, &orderData.PostId, &orderData.HouseId, &orderData.UserId, &orderData.BuyUserId, &orderData.State, &orderData.Price)
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
			Data: orderData,
		})
	})
}
