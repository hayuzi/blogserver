package v1

import "github.com/hayuzi/blogserver/internal/model"

// ArticleListReq 文章列表
type ArticleListReq struct {
	Id            int    `json:"id" form:"id" binding:"min=0"`
	Title         string `json:"title" form:"title" binding:"max=255"`
	TagId         int    `json:"tagId" form:"tagId" binding:"min=0"`
	ArticleStatus int    `json:"articleStatus" form:"articleStatus" binding:"min=0"`
	PageNum       int    `json:"pageNum" form:"pageNum"`
	PageSize      int    `json:"pageSize" form:"pageSize"`
}
type ArticleListRes struct {
	Lists []model.Article `json:"lists"`
	Total int64           `json:"total"`
}

// ArticleCreateReq 文章创建
type ArticleCreateReq struct {
	Title         string `json:"title" form:"title" binding:"required,min=1,max=255"`
	Sketch        string `json:"sketch" form:"sketch" binding:"min=0,max=255"`
	Content       string `json:"content" form:"content" binding:"required,min=1,max=65535"`
	TagId         int    `json:"tagId" form:"tagId" binding:"required,min=1"`
	Weight        int    `json:"weight" form:"weight" binding:"min=0,max=100"`
	ArticleStatus int    `json:"articleStatus" form:"articleStatus" binding:"required,min=1,max=2"`
}
type ArticleCreateRes struct {
	Id int `json:"id"`
}

// ArticleUpdateReq 编辑
type ArticleUpdateReq struct {
	Id            int    `json:"id" form:"id"` // 从路由获取注入
	Title         string `json:"title" form:"title" binding:"required,min=1,max=255"`
	Sketch        string `json:"sketch" form:"sketch" binding:"min=0,max=255"`
	Content       string `json:"content" form:"content" binding:"required,min=1,max=65535"`
	TagId         int    `json:"tagId" form:"tagId" binding:"required,min=1"`
	Weight        int    `json:"weight" form:"weight" binding:"min=0,max=100"`
	ArticleStatus int    `json:"articleStatus" form:"articleStatus" binding:"required,min=1,max=2"`
}
type ArticleUpdateRes struct {
	Id int `json:"id"`
}

// ArticleDeleteReq 文章删除
type ArticleDeleteReq struct {
	Id int `json:"id" form:"id"`
}
type ArticleDeleteRes struct {
	Id int `json:"id"`
}
