package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/hayuzi/blogserver/internal/controller/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	articleV1 := v1.NewArticle()
	commentV1 := v1.NewComment()
	tagV1 := v1.NewTag()

	// 通用接口
	// 业务接口v1
	apiV1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiV1.GET("/tags", tagV1.List)
		apiV1.GET("/tag", tagV1.Get)
		//获取所有标签
		apiV1.GET("/tags/all", tagV1.All)

		//获取文章列表
		apiV1.GET("/articles", articleV1.List)
		//获取指定文章
		apiV1.GET("/articles/:id", articleV1.Get)

		//获取评论列表
		apiV1.GET("/comments", commentV1.List)
		//获取评论
		apiV1.GET("/comments/:id", commentV1.Get)
	}

	return r
}
