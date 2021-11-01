package v1

import "github.com/hayuzi/blogserver/internal/model"

// TagListReq  标签列表
type TagListReq struct {
	Id        int    `json:"id" form:"id" binding:"min=0"`
	TagName   string `json:"tagName" form:"tagName" binding:"max=100"`
	TagStatus int    `json:"tagStatus" form:"tagStatus" binding:"min=0"`
	PageNum   int    `json:"pageNum" form:"pageNum"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
}
type TagListRes struct {
	Lists []model.Tag `json:"lists"`
	Total int64       `json:"total"`
}

// TagAllReq  标签全列表
type TagAllReq struct{}
type TagAllRes struct {
	Lists []model.Tag `json:"lists"`
}

// TagCreateReq 标签创建
type TagCreateReq struct {
	TagName   string `json:"tagName" form:"tagName" binding:"required,min=1,max=64"`
	Weight    int    `json:"weight" form:"weight" binding:"min=0,max=100"`
	TagStatus int    `json:"tagStatus" form:"tagStatus" binding:"required,min=1,max=2"`
}
type TagCreateRes struct {
	Id int `json:"id"`
}

// TagUpdateReq 标签编辑
type TagUpdateReq struct {
	Id        int    `json:"id" form:"id"` // 从路由获取注入
	TagName   string `json:"tagName" form:"tagName" binding:"required,min=1,max=64"`
	Weight    int    `json:"weight" form:"weight" binding:"min=0,max=100"`
	TagStatus int    `json:"tagStatus" form:"tagStatus" binding:"required,min=1,max=2"`
}
type TagUpdateRes struct {
	Id int `json:"id"`
}

// TagDeleteReq 标签删除
type TagDeleteReq struct {
	Id int `json:"id" form:"id"`
}
type TagDeleteRes struct {
	Id int `json:"id"`
}
