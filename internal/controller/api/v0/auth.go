package v0

import (
	"github.com/gin-gonic/gin"
	fmtApiV0 "github.com/hayuzi/blogserver/internal/fmtter/api/v0"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/internal/service"
	"github.com/hayuzi/blogserver/pkg/app"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

type Auth struct{}

func NewAuth() Auth {
	return Auth{}
}

func (t Auth) AuthRegister(c *gin.Context) {
	req := fmtApiV0.AuthRegisterReq{}
	res := fmtApiV0.AuthRegisterRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c)
	cusErr := svc.AuthRegister(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}

func (t Auth) AuthLogin(c *gin.Context) {
	req := fmtApiV0.AuthLoginReq{}
	res := fmtApiV0.AuthLoginRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c)
	cusErr := svc.AuthLogin(model.UserTypeUser, &req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}

func (t Auth) AuthAdminLogin(c *gin.Context) {
	req := fmtApiV0.AuthLoginReq{}
	res := fmtApiV0.AuthLoginRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c)
	cusErr := svc.AuthLogin(model.UserTypeAdmin, &req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}
