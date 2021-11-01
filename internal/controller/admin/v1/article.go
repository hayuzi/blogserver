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

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (t Article) Get(c *gin.Context) {
	res := model.Article{}
	response := app.NewResponse(c)
	id, _ := strconv.Atoi(c.Param("id"))
	svc := service.New(c.Request.Context())
	cusErr := svc.ArticleDetailAdmin(id, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}
func (t Article) List(c *gin.Context) {
	req := fmtAdminV1.ArticleListReq{}
	res := fmtAdminV1.ArticleListRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	cusErr := svc.ArticleListAdmin(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponseList(res.Lists, res.Total)
	return
}

func (t Article) Create(c *gin.Context) {
	req := fmtAdminV1.ArticleCreateReq{}
	res := fmtAdminV1.ArticleCreateRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	cusErr := svc.ArticleCreateAdmin(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}

func (t Article) Update(c *gin.Context) {
	req := fmtAdminV1.ArticleUpdateReq{}
	res := fmtAdminV1.ArticleUpdateRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	req.Id, _ = strconv.Atoi(c.Param("id"))
	svc := service.New(c.Request.Context())
	cusErr := svc.ArticleUpdateAdmin(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
}

func (t Article) Delete(c *gin.Context) {
	req := fmtAdminV1.ArticleDeleteReq{}
	res := fmtAdminV1.ArticleDeleteRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	req.Id, _ = strconv.Atoi(c.Param("id"))
	svc := service.New(c.Request.Context())
	cusErr := svc.ArticleDeleteAdmin(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}
