package model

type Order struct {
	Id        int `json:"id"`
	PostId    int `json:"post_id" binding:"required"`
	HouseId   int `json:"house_id" binding:"required"`
	UserId    int `json:"user_id" binding:"required"`
	BuyUserId int `json:"buy_user_id" binding:"required"`
	State     int `json:"state" binding:"required"`
	Price     int `json:"price" binding:"required"`
}

// 完成订单所需要传递的参数
type ReqCompleteOrder struct {
	OrderId int `json:"order_id" binding:"required"`
	PostId  int `json:"post_id" binding:"required"`
	HouseId int `json:"house_id" binding:"required"`
}

// func NewOrder(order *Order) *Order {
// 	return &Order{
// 		PostId:    order.PostId,
// 		HouseId:   order.PostId,
// 		UserId:    order.UserId,
// 		BuyUserId: order.BuyUserId,
// 		State:     order.State,
// 		Price:     order.Price,
// 	}
// }
