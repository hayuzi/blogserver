package service

import (
	fmtAdminV1 "github.com/hayuzi/blogserver/internal/fmtter/admin/v1"
	fmtApiV1 "github.com/hayuzi/blogserver/internal/fmtter/api/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func (svc *Service) ArticleList(req *fmtApiV1.ArticleListReq, res *fmtApiV1.ArticleListRes) *errcode.Error {
	err := svc.dao.ArticlePaginatedList(svc.ctx, req, res)
	if err != nil {
		return errcode.ArticleListFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) ArticleListAdmin(req *fmtAdminV1.ArticleListReq, res *fmtAdminV1.ArticleListRes) *errcode.Error {
	err := svc.dao.ArticlePaginatedListAdmin(svc.ctx, req, res)
	if err != nil {
		return errcode.ArticleListFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) ArticleCreateAdmin(req *fmtAdminV1.ArticleCreateReq, res *fmtAdminV1.ArticleCreateRes) *errcode.Error {
	err := svc.dao.ArticleCreateAdmin(svc.ctx, req, res)
	if err != nil {
		return errcode.ArticleCreateFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) ArticleUpdateAdmin(req *fmtAdminV1.ArticleUpdateReq, res *fmtAdminV1.ArticleUpdateRes) *errcode.Error {
	err := svc.dao.ArticleUpdateAdmin(svc.ctx, req, res)
	if err != nil {
		return errcode.ArticleUpdateFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) ArticleDeleteAdmin(req *fmtAdminV1.ArticleDeleteReq, res *fmtAdminV1.ArticleDeleteRes) *errcode.Error {
	err := svc.dao.ArticleDeleteAdmin(svc.ctx, req, res)
	if err != nil {
		return errcode.ArticleDeleteFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) ArticleDetail(id int, res *model.Article) *errcode.Error {
	err := svc.dao.ArticleDetail(svc.ctx, id, res, true)
	if res.ArticleStatus != model.ArticleStatusPublished {
		return errcode.ArticleDetailFail.WithDetails("文章草稿不可以访问")
	}
	if err != nil {
		return errcode.ArticleDetailFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) ArticleDetailAdmin(id int, res *model.Article) *errcode.Error {
	err := svc.dao.ArticleDetail(svc.ctx, id, res, true)
	if err != nil {
		return errcode.ArticleDetailFail.WithDetails(err.Error())
	}
	return nil
}
