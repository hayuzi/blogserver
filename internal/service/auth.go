package service

import (
	fmtApiV0 "github.com/hayuzi/blogserver/internal/fmtter/api/v0"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func (svc *Service) AuthRegister(req *fmtApiV0.AuthRegisterReq, res *fmtApiV0.AuthRegisterRes) *errcode.Error {
	// 用户名是否被占用
	exists, err := svc.dao.UserExistsByUsername(svc.ctx, req.Username, 0)
	if err != nil {
		return errcode.AuthRegisterFail.WithDetails(err.Error())
	}
	if exists {
		return errcode.AuthUsernameExists
	}
	// 创建用户
	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Pwd:      app.EncodePwd(req.Pwd),
		UserType: model.UserTypeUser,
	}
	err = svc.dao.UserCreate(svc.ctx, &user)
	if err != nil {
		return errcode.AuthRegisterFail.WithDetails(err.Error())
	}
	// 注册即可登陆
	res.Token, err = app.GenerateToken(user.Id, user.Username, user.UserType)
	if err != nil {
		return errcode.AuthTokenGenerateFail.WithDetails(err.Error())
	}
	res.Id = user.Id
	res.Username = user.Username
	res.Email = user.Email
	res.UserType = user.UserType
	return nil
}

func (svc *Service) AuthLogin(userType int, req *fmtApiV0.AuthLoginReq, res *fmtApiV0.AuthLoginRes) *errcode.Error {
	user := model.User{}
	err := svc.dao.UserByUsername(svc.ctx, req.Username, &user)
	if err != nil {
		return errcode.AuthUserNotExists.WithDetails(err.Error())
	}
	// TODO@yuzi 增加登陆密码错误次数限制（ 可以选择缓存Redis、单机可以使用进程内缓存（重启数据会丢失） ）
	if app.EncodePwd(req.Pwd) != user.Pwd {
		return errcode.AuthPwdNotCorrect
	}
	if userType == model.UserTypeAdmin && userType != user.UserType {
		return errcode.UnauthorizedUserNotAdmin
	}
	res.Token, err = app.GenerateToken(user.Id, user.Username, user.UserType)
	if err != nil {
		return errcode.AuthTokenGenerateFail.WithDetails(err.Error())
	}
	res.Id = user.Id
	res.Username = user.Username
	res.Email = user.Email
	res.UserType = user.UserType
	return nil
}
