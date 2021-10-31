package v1

import "github.com/hayuzi/blogserver/internal/model"

// UserListReq 列表
type UserListReq struct {
	Id       int    `json:"id" form:"id" binding:"min=0"`
	Username string `json:"title" form:"title" binding:"max=64"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}
type UserListRes struct {
	Lists []model.User `json:"lists"`
	Total int64        `json:"total"`
}

// UserCreateReq 创建
type UserCreateReq struct {
	Username string `json:"username" form:"username" binding:"required,min=1,max=64"`
	Email    string `json:"email" form:"email" binding:"required,min=1,max=64"`
	Pwd      string `json:"pwd" form:"pwd" binding:"required,min=1,max=64"`
}
type UserCreateRes struct {
	Id int `json:"id"`
}

// UserUpdateReq 编辑
type UserUpdateReq struct {
	Id       int    `json:"id" form:"id"` // 从路由获取注入
	Username string `json:"username" form:"username" binding:"required,min=1,max=64"`
	Email    string `json:"email" form:"email" binding:"required,min=1,max=64"`
	Pwd      string `json:"pwd" form:"pwd" binding:"required,min=1,max=64"`
}
type UserUpdateRes struct {
	Id int `json:"id"`
}

// UserDeleteReq 删除
type UserDeleteReq struct {
	Id int `json:"id" form:"id"`
}
type UserDeleteRes struct {
	Id int `json:"id"`
}
