package model

type Callback struct {
	Attach          int `json:"attach"`
	CreateTime      int `json:"create_time" `
	Fee             int `json:"fee"`
	MerchanOrderNum int `json:"merchant_order_num"`
	PointFee        int `json:"point_fee"`
	State           int `json:"state"`
	UpdateTime      int `json:"update_time"`
}
