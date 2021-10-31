package v1

import (
	"github.com/gin-gonic/gin"
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
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
	cusErr := svc.ArticleDetail(c, id, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}
func (t Article) List(c *gin.Context) {
	req := fmtV1.ArticleListReq{}
	res := fmtV1.ArticleListRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	cusErr := svc.ArticleList(c, &req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponseList(res.Lists, res.Total)
	return
}

func (t Article) Create(c *gin.Context) {
	req := fmtV1.ArticleCreateReq{}
	res := fmtV1.ArticleCreateRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	cusErr := svc.ArticleCreate(c, &req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}

func (t Article) Update(c *gin.Context) {
	req := fmtV1.ArticleUpdateReq{}
	res := fmtV1.ArticleUpdateRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	req.Id, _ = strconv.Atoi(c.Param("id"))
	svc := service.New(c.Request.Context())
	cusErr := svc.ArticleUpdate(c, &req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
}

func (t Article) Delete(c *gin.Context) {
	req := fmtV1.ArticleDeleteReq{}
	res := fmtV1.ArticleDeleteRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	req.Id, _ = strconv.Atoi(c.Param("id"))
	svc := service.New(c.Request.Context())
	cusErr := svc.ArticleDelete(c, &req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}
