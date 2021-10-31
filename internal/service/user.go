package service

import (
	"github.com/gin-gonic/gin"
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func (svc *Service) UserList(c *gin.Context, req *fmtV1.UserListReq, res *fmtV1.UserListRes) *errcode.Error {
	err := svc.dao.UserPaginatedList(c.Request.Context(), req, res)
	if err != nil {
		return errcode.UserListFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) UserCreate(c *gin.Context, req *fmtV1.UserCreateReq, res *fmtV1.UserCreateRes) *errcode.Error {
	err := svc.dao.UserCreate(c.Request.Context(), req, res)
	if err != nil {
		return errcode.UserCreateFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) UserUpdate(c *gin.Context, req *fmtV1.UserUpdateReq, res *fmtV1.UserUpdateRes) *errcode.Error {
	err := svc.dao.UserUpdate(c.Request.Context(), req, res)
	if err != nil {
		return errcode.UserUpdateFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) UserDelete(c *gin.Context, req *fmtV1.UserDeleteReq, res *fmtV1.UserDeleteRes) *errcode.Error {
	err := svc.dao.UserDelete(c.Request.Context(), req, res)
	if err != nil {
		return errcode.UserDeleteFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) UserDetail(c *gin.Context, id int, res *model.User) *errcode.Error {
	err := svc.dao.UserDetail(c.Request.Context(), id, res)
	if err != nil {
		return errcode.UserDetailFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}
