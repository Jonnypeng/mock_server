package model

// 定义数据的返回结构体
type ResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
