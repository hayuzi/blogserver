package v1

import (
	"github.com/gin-gonic/gin"
	fmtApiV1 "github.com/hayuzi/blogserver/internal/fmtter/api/v1"
	"github.com/hayuzi/blogserver/internal/service"
	"github.com/hayuzi/blogserver/pkg/app"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (t User) ChangePwd(c *gin.Context) {
	req := fmtApiV1.UserChangePwdReq{}
	res := fmtApiV1.UserChangePwdRes{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if valid == true {
		response.ToResponseError(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c)
	cusErr := svc.UserChangePwd(&req, &res)
	if cusErr != nil {
		response.ToResponseError(cusErr)
		return
	}
	response.ToResponse(res)
	return
}
