package v1

import "github.com/hayuzi/blogserver/internal/model"

// ArticleListReq 文章列表
type ArticleListReq struct {
	Id       int    `json:"id" form:"id" binding:"min=0"`
	Q        string `json:"q" form:"q" binding:"max=255"`
	TagId    int    `json:"tagId" form:"tagId" binding:"min=0"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}
type ArticleListRes struct {
	Lists []model.Article `json:"lists"`
	Total int64           `json:"total"`
}
