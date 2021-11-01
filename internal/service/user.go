package service

import (
	fmtAdminV1 "github.com/hayuzi/blogserver/internal/fmtter/admin/v1"
	fmtApiV1 "github.com/hayuzi/blogserver/internal/fmtter/api/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func (svc *Service) UserListAdmin(req *fmtAdminV1.UserListReq, res *fmtAdminV1.UserListRes) *errcode.Error {
	err := svc.dao.UserPaginatedListAdmin(svc.ctx, req, res)
	if err != nil {
		return errcode.UserListFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) UserDeleteAdmin(req *fmtAdminV1.UserDeleteReq, res *fmtAdminV1.UserDeleteRes) *errcode.Error {
	err := svc.dao.UserDelete(svc.ctx, req.Id)
	if err != nil {
		return errcode.UserDeleteFail.WithDetails(err.Error())
	}
	res.Id = req.Id
	return nil
}

func (svc *Service) UserDetail(id int, res *model.User) *errcode.Error {
	err := svc.dao.UserDetail(svc.ctx, id, res)
	if err != nil {
		return errcode.UserDetailFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) UserChangePwd(req *fmtApiV1.UserChangePwdReq, res *fmtApiV1.UserChangePwdRes) *errcode.Error {
	if req.Confirm != req.Pwd {
		return errcode.UserPwdConfirmFail
	}
	loginUser := app.GetLoginUser(svc.ctx)
	userId := loginUser.Id
	userInfo := model.User{}
	err := svc.dao.UserDetail(svc.ctx, userId, &userInfo)
	if err != nil {
		return errcode.UserDetailFail.WithDetails(err.Error())
	}
	encodedOldPwd := app.EncodePwd(req.Pwd)
	if encodedOldPwd != userInfo.Pwd {
		return errcode.UserOldPwdError
	}

	encodedNewPwd := app.EncodePwd(req.Password)
	err = svc.dao.UserChangePwd(svc.ctx, userId, encodedNewPwd)
	if err != nil {
		return errcode.UserDetailFail.WithDetails(err.Error())
	}
	res.Id = loginUser.Id
	return nil
}
