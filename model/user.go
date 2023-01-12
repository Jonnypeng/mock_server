package model

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	UserImg  string `json:"user_img"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserInfo struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	UserImg  string `json:"user_img"`
	Phone    string `json:"phone" binding:"required"`
}

func NewUserInfo(user User) *UserInfo {
	return &UserInfo{
		Id:       user.Id,
		UserName: user.UserName,
		Phone:    user.Phone,
		UserImg:  user.UserImg,
	}
}
