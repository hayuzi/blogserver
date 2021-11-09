package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/global"
	adminV1 "github.com/hayuzi/blogserver/internal/controller/admin/v1"
	v0 "github.com/hayuzi/blogserver/internal/controller/api/v0"
	v1 "github.com/hayuzi/blogserver/internal/controller/api/v1"
	"github.com/hayuzi/blogserver/internal/midddleware"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(midddleware.AccessLog())

	// 限流
	//limiter := ratelimiter.NewMethodLimiter()
	//limiter.AddBuckets(ratelimiter.LimiterBucketRule{
	//	Key:          "/api/v1/tags",
	//	FillInterval: time.Second * 10,
	//	Capacity:     10,
	//	Quantum:      1,
	//})
	//r.Use(midddleware.RateLimiter(limiter))

	r.Use(midddleware.Translations())
	r.Use(midddleware.Cors())
	r.Use(midddleware.JWTInjectClaims())

	v0Auth := v0.NewAuth()
	v0Upload := v0.NewUpload()

	// 静态文件访问
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// 通用接口
	apiV0 := r.Group("/api/v0")
	{
		apiV0.POST("/register", v0Auth.AuthRegister)
		apiV0.GET("/auth", v0Auth.AuthLogin)
		apiV0.GET("/admin-auth", v0Auth.AuthAdminLogin)
		apiV0.POST("/upload", v0Upload.UploadFile)
	}

	// 业务接口v1
	v1Article := v1.NewArticle()
	v1Comment := v1.NewComment()
	v1Tag := v1.NewTag()
	v1User := v1.NewUser()

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tags", v1Tag.List)    //获取标签列表
		apiV1.GET("/tags/all", v1Tag.All) //获取所有标签
		apiV1.GET("/tags/:id", v1Tag.Get) //获取标签详情

		apiV1.GET("/articles", v1Article.List)    //获取文章列表
		apiV1.GET("/articles/:id", v1Article.Get) //获取指定文章

		apiV1.GET("/comments", v1Comment.List)    //获取评论列表
		apiV1.GET("/comments/:id", v1Comment.Get) //获取评论
	}

	// 业务接口v1，需要登陆
	apiV1Auth := r.Group("/api/v1/auth")
	apiV1Auth.Use(midddleware.JWT())
	{
		apiV1Auth.POST("/comments", v1Comment.Create)  // 添加评论
		apiV1Auth.PUT("/change-pwd", v1User.ChangePwd) // 修改密码
	}

	// 控制器获取
	adminV1Article := adminV1.NewArticle()
	adminV1Comment := adminV1.NewComment()
	adminV1Common := adminV1.NewCommon()
	adminV1Tag := adminV1.NewTag()
	adminV1User := adminV1.NewUser()
	// 管理后台接口v1
	apiAdminV1 := r.Group("/admin/v1")
	apiAdminV1.Use(midddleware.JWTAdmin())
	{
		apiAdminV1.GET("/dashboard", adminV1Common.Dashboard) //获取调色盘信息
		apiAdminV1.GET("/tags", adminV1Tag.List)              //获取标签列表
		apiAdminV1.POST("/tags", adminV1Tag.Create)           //新建标签
		apiAdminV1.GET("/tags/all", adminV1Tag.All)           //获取所有标签
		apiAdminV1.PUT("/tags/:id", adminV1Tag.Update)        //更新指定标签
		apiAdminV1.DELETE("/tags/:id", adminV1Tag.Delete)     //删除指定标签

		apiAdminV1.GET("/articles", adminV1Article.List)          //获取文章列表
		apiAdminV1.GET("/articles/:id", adminV1Article.Get)       //获取指定文章
		apiAdminV1.POST("/articles", adminV1Article.Create)       //新建文章
		apiAdminV1.PUT("/articles/:id", adminV1Article.Update)    //更新指定文章
		apiAdminV1.DELETE("/articles/:id", adminV1Article.Delete) //删除指定文章

		apiAdminV1.GET("/comments", adminV1Comment.List)          //获取评论列表
		apiAdminV1.DELETE("/comments/:id", adminV1Comment.Delete) //删除评论

		apiAdminV1.GET("/users", adminV1User.List)          //获取用户列表
		apiAdminV1.DELETE("/users/:id", adminV1User.Delete) //删除用户
	}

	return r
}
