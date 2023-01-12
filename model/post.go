package model

type Post struct {
	Id            int    `json:"id"`
	LicenseNumber string `json:"license_number" binding:"required"`
	UserId        int    `json:"user_id" binding:"required"`
	PostType      int    `json:"post_type" binding:"required"`
	CreateTime    string `json:"create_time"`
	Price         int    `json:"price" binding:"required"`
	State         int    `json:"state"`
	Title         string `json:"title"`
	Imgs          string `json:"imgs"`
	HouseType     string `json:"house_type"`
	HouseId       int    `json:"house_id"`
}

type ReqPostList struct {
	Cursor int `json:"cursor" binding:"required"`
	Limit  int `json:"limit" bingding:"required"`
}

// func NewPost(post *Post) *Post {
// 	return &Post{
// 		HouseId:    post.HouseId,
// 		UserId:     post.UserId,
// 		PostType:   post.PostType,
// 		CreateTime: post.CreateTime,
// 		Price:      post.Price,
// 		State:      post.State,
// 	}
// }
