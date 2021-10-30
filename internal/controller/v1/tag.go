package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/global"
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
	"github.com/hayuzi/blogserver/internal/service"
	"github.com/hayuzi/blogserver/pkg/app"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	data := map[string]interface{}{
		"test": 1,
	}
	app.NewResponse(c).ToResponse(data)
	return
}
func (t Tag) List(c *gin.Context) {
	req := fmtV1.TagListReq{}
	res := fmtV1.TagListRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	cusErr := svc.TagList(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponseList(res.Lists, res.Total)
	return
}
func (t Tag) All(c *gin.Context) {
	app.NewResponse(c).ToResponseError(errcode.ServerError)
	return
}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
