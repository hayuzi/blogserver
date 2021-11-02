package v1

import "github.com/hayuzi/blogserver/internal/model"

// CommentListReq 列表
type CommentListReq struct {
	Id            int `json:"id" form:"id" binding:"min=0"`
	ArticleId     int `json:"articleId" form:"articleId" binding:"min=0"`
	UserId        int `json:"userId" form:"userId" binding:"min=0"`
	CommentStatus int `json:"commentStatus" form:"commentStatus" binding:"min=0"`
	PageNum       int `json:"pageNum" form:"pageNum"`
	PageSize      int `json:"pageSize" form:"pageSize"`
}
type CommentListRes struct {
	Lists []model.Comment `json:"lists"`
	Total int64           `json:"total"`
}

// CommentCreateReq 创建
type CommentCreateReq struct {
	UserId        int    `json:"userId" form:"userId"` // 登陆状态中获取
	MentionUserId int    `json:"mentionUserId" form:"mentionUserId" binding:"min=0"`
	ArticleId     int    `json:"articleId" form:"articleId" binding:"required,min=0"`
	Content       string `json:"content" form:"content" binding:"required,min=1,max=1024"`
}
type CommentCreateRes struct {
	Id int `json:"id"`
}

// CommentUpdateReq 编辑
type CommentUpdateReq struct {
	Id            int    `json:"id" form:"id"` // 从路由获取注入
	ArticleId     int    `json:"articleId" form:"articleId" binding:"required,min=0"`
	UserId        int    `json:"userId" form:"userId" binding:"required,min=0"`
	MentionUserId int    `json:"mentionUserId" form:"mentionUserId" binding:"required,min=0"`
	Content       string `json:"content" form:"title" binding:"required,min=1,max=1024"`
	CommentStatus int    `json:"commentStatus" form:"commentStatus" binding:"required,min=1,max=2"`
}
type CommentUpdateRes struct {
	Id int `json:"id"`
}

// CommentDeleteReq 删除
type CommentDeleteReq struct {
	Id int `json:"id" form:"id"`
}
type CommentDeleteRes struct {
	Id int `json:"id"`
}
