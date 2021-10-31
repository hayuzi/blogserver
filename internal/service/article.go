package service

import (
	"github.com/gin-gonic/gin"
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func (svc *Service) ArticleList(c *gin.Context, req *fmtV1.ArticleListReq, res *fmtV1.ArticleListRes) *errcode.Error {
	err := svc.dao.ArticlePaginatedList(c.Request.Context(), req, res)
	if err != nil {
		return errcode.ArticleListFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) ArticleCreate(c *gin.Context, req *fmtV1.ArticleCreateReq, res *fmtV1.ArticleCreateRes) *errcode.Error {
	err := svc.dao.ArticleCreate(c.Request.Context(), req, res)
	if err != nil {
		return errcode.ArticleCreateFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) ArticleUpdate(c *gin.Context, req *fmtV1.ArticleUpdateReq, res *fmtV1.ArticleUpdateRes) *errcode.Error {
	err := svc.dao.ArticleUpdate(c.Request.Context(), req, res)
	if err != nil {
		return errcode.ArticleUpdateFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) ArticleDelete(c *gin.Context, req *fmtV1.ArticleDeleteReq, res *fmtV1.ArticleDeleteRes) *errcode.Error {
	err := svc.dao.ArticleDelete(c.Request.Context(), req, res)
	if err != nil {
		return errcode.ArticleDeleteFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}

func (svc *Service) ArticleDetail(c *gin.Context, id int, res *model.Article) *errcode.Error {
	err := svc.dao.ArticleDetail(c.Request.Context(), id, res)
	if err != nil {
		return errcode.ArticleDetailFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}
