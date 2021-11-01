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

type Comment struct{}

func NewComment() Comment {
	return Comment{}
}

func (t Comment) Get(c *gin.Context) {
	res := model.Comment{}
	response := app.NewResponse(c)
	id, _ := strconv.Atoi(c.Param("id"))
	svc := service.New(c.Request.Context())
	cusErr := svc.CommentDetail(id, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}
func (t Comment) List(c *gin.Context) {
	req := fmtApiV1.CommentListReq{}
	res := fmtApiV1.CommentListRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	cusErr := svc.CommentList(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponseList(res.Lists, res.Total)
	return
}

func (t Comment) Create(c *gin.Context) {
	req := fmtApiV1.CommentCreateReq{}
	res := fmtApiV1.CommentCreateRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	cusErr := svc.CommentCreate(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}
