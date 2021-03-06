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
