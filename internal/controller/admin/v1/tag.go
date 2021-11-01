package v1

import (
	"github.com/gin-gonic/gin"
	fmtAdminV1 "github.com/hayuzi/blogserver/internal/fmtter/admin/v1"
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
	svc := service.New(c.Request.Context())
	cusErr := svc.TagDetail(id, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}

func (t Tag) List(c *gin.Context) {
	req := fmtAdminV1.TagListReq{}
	res := fmtAdminV1.TagListRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	cusErr := svc.TagListAdmin(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponseList(res.Lists, res.Total)
	return
}

func (t Tag) All(c *gin.Context) {
	req := fmtAdminV1.TagAllReq{}
	res := fmtAdminV1.TagAllRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	cusErr := svc.TagAllAdmin(&res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponseList(res.Lists, int64(len(res.Lists)))
	return
}

func (t Tag) Create(c *gin.Context) {
	req := fmtAdminV1.TagCreateReq{}
	res := fmtAdminV1.TagCreateRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	cusErr := svc.TagCreateAdmin(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}

func (t Tag) Update(c *gin.Context) {
	req := fmtAdminV1.TagUpdateReq{}
	res := fmtAdminV1.TagUpdateRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	req.Id, _ = strconv.Atoi(c.Param("id"))
	svc := service.New(c.Request.Context())
	cusErr := svc.TagUpdateAdmin(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}

func (t Tag) Delete(c *gin.Context) {
	req := fmtAdminV1.TagDeleteReq{}
	res := fmtAdminV1.TagDeleteRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	req.Id, _ = strconv.Atoi(c.Param("id"))
	svc := service.New(c.Request.Context())
	cusErr := svc.TagDeleteAdmin(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}
