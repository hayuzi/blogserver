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

type User struct{}

func NewUser() User {
	return User{}
}

func (t User) Get(c *gin.Context) {
	res := model.User{}
	response := app.NewResponse(c)
	id, _ := strconv.Atoi(c.Param("id"))
	svc := service.New(c)
	cusErr := svc.UserDetail(id, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}
func (t User) List(c *gin.Context) {
	req := fmtAdminV1.UserListReq{}
	res := fmtAdminV1.UserListRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c)
	cusErr := svc.UserListAdmin(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponseList(res.Lists, res.Total)
	return
}

func (t User) Delete(c *gin.Context) {
	req := fmtAdminV1.UserDeleteReq{}
	res := fmtAdminV1.UserDeleteRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	req.Id, _ = strconv.Atoi(c.Param("id"))
	svc := service.New(c)
	cusErr := svc.UserDeleteAdmin(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}
