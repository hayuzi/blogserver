package service

import (
	"github.com/gin-gonic/gin"
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func (svc *Service) TagList(c *gin.Context, req *fmtV1.TagListReq, res *fmtV1.TagListRes) *errcode.Error {
	err := svc.dao.TagPaginatedList(c.Request.Context(), req, res)
	if err != nil {
		return errcode.TagListFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) TagAll(c *gin.Context, res *fmtV1.TagAllRes) *errcode.Error {
	err := svc.dao.TagAll(c.Request.Context(), res)
	if err != nil {
		return errcode.TagListFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) TagCreate(c *gin.Context, req *fmtV1.TagCreateReq, res *fmtV1.TagCreateRes) *errcode.Error {
	err := svc.dao.TagCreate(c.Request.Context(), req, res)
	if err != nil {
		return errcode.TagCreateFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) TagUpdate(c *gin.Context, req *fmtV1.TagUpdateReq, res *fmtV1.TagUpdateRes) *errcode.Error {
	err := svc.dao.TagUpdate(c.Request.Context(), req, res)
	if err != nil {
		return errcode.TagUpdateFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) TagDelete(c *gin.Context, req *fmtV1.TagDeleteReq, res *fmtV1.TagDeleteRes) *errcode.Error {
	err := svc.dao.TagDelete(c.Request.Context(), req, res)
	if err != nil {
		return errcode.TagDeleteFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) TagDetail(c *gin.Context, id int, res *model.Tag) *errcode.Error {
	err := svc.dao.TagDetail(c.Request.Context(), id, res)
	if err != nil {
		return errcode.TagDetailFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}
