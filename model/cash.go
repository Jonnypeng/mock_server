package model

type Callback struct {
	Attach          int `json:"attach" binding:"required"`
	CreateTime      int `json:"create_time" binding:"required"`
	Fee             int `json:"fee" binding:"required"`
	MerchanOrderNum int `json:"merchant_order_num" binding:"required"`
	PointFee        int `json:"point_fee" binding:"required"`
	State           int `json:"state" binding:"required"`
	UpdateTime      int `json:"update_time" binding:"required"`
}
