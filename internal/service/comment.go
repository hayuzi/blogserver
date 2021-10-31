package service

import (
	"github.com/gin-gonic/gin"
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func (svc *Service) CommentList(c *gin.Context, req *fmtV1.CommentListReq, res *fmtV1.CommentListRes) *errcode.Error {
	err := svc.dao.CommentPaginatedList(c.Request.Context(), req, res)
	if err != nil {
		return errcode.CommentListFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) CommentCreate(c *gin.Context, req *fmtV1.CommentCreateReq, res *fmtV1.CommentCreateRes) *errcode.Error {
	err := svc.dao.CommentCreate(c.Request.Context(), req, res)
	if err != nil {
		return errcode.CommentCreateFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) CommentUpdate(c *gin.Context, req *fmtV1.CommentUpdateReq, res *fmtV1.CommentUpdateRes) *errcode.Error {
	err := svc.dao.CommentUpdate(c.Request.Context(), req, res)
	if err != nil {
		return errcode.CommentUpdateFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) CommentDelete(c *gin.Context, req *fmtV1.CommentDeleteReq, res *fmtV1.CommentDeleteRes) *errcode.Error {
	err := svc.dao.CommentDelete(c.Request.Context(), req, res)
	if err != nil {
		return errcode.CommentDeleteFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) CommentDetail(c *gin.Context, id int, res *model.Comment) *errcode.Error {
	err := svc.dao.CommentDetail(c.Request.Context(), id, res)
	if err != nil {
		return errcode.CommentDetailFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}
