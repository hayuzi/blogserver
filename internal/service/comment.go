package service

import (
	fmtAdminV1 "github.com/hayuzi/blogserver/internal/fmtter/admin/v1"
	fmtApiV1 "github.com/hayuzi/blogserver/internal/fmtter/api/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func (svc *Service) CommentList(req *fmtApiV1.CommentListReq, res *fmtApiV1.CommentListRes) *errcode.Error {
	err := svc.dao.CommentPaginatedList(svc.ctx, req, res)
	if err != nil {
		return errcode.CommentListFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) CommentListAdmin(req *fmtAdminV1.CommentListReq, res *fmtAdminV1.CommentListRes) *errcode.Error {
	err := svc.dao.CommentPaginatedListAdmin(svc.ctx, req, res)
	if err != nil {
		return errcode.CommentListFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) CommentCreate(req *fmtApiV1.CommentCreateReq, res *fmtApiV1.CommentCreateRes) *errcode.Error {
	loginUser := app.GetLoginUser(svc.ctx)
	req.UserId = loginUser.Id
	err := svc.dao.CommentCreate(svc.ctx, req, res)
	if err != nil {
		return errcode.CommentCreateFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) CommentUpdate(req *fmtApiV1.CommentUpdateReq, res *fmtApiV1.CommentUpdateRes) *errcode.Error {
	err := svc.dao.CommentUpdate(svc.ctx, req, res)
	if err != nil {
		return errcode.CommentUpdateFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) CommentDeleteAdmin(req *fmtAdminV1.CommentDeleteReq, res *fmtAdminV1.CommentDeleteRes) *errcode.Error {
	err := svc.dao.CommentDeleteAdmin(svc.ctx, req, res)
	if err != nil {
		return errcode.CommentDeleteFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) CommentDetail(id int, res *model.Comment) *errcode.Error {
	err := svc.dao.CommentDetail(svc.ctx, id, res)
	if err != nil {
		return errcode.CommentDetailFail.WithDetails(err.Error())
	}
	return nil
}
