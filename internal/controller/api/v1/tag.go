package v1

import (
	"github.com/gin-gonic/gin"
	fmtApiV1 "github.com/hayuzi/blogserver/internal/fmtter/api/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/internal/service"
	"github.com/hayuzi/blogserver/pkg/app"
	"github.com/hayuzi/blogserver/pkg/errcode"
	"strconv"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	res := model.Tag{}
	response := app.NewResponse(c)
	id, _ := strconv.Atoi(c.Param("id"))
	svc := service.New(c)
	cusErr := svc.TagDetail(id, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}

func (t Tag) List(c *gin.Context) {
	req := fmtApiV1.TagListReq{}
	res := fmtApiV1.TagListRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c)
	cusErr := svc.TagList(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponseList(res.Lists, res.Total)
	return
}

func (t Tag) All(c *gin.Context) {
	req := fmtApiV1.TagAllReq{}
	res := fmtApiV1.TagAllRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c)
	cusErr := svc.TagAll(&res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponseList(res.Lists, int64(len(res.Lists)))
	return
}
