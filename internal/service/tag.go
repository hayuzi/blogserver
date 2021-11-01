package service

import (
	fmtAdminV1 "github.com/hayuzi/blogserver/internal/fmtter/admin/v1"
	fmtApiV1 "github.com/hayuzi/blogserver/internal/fmtter/api/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func (svc *Service) TagList(req *fmtApiV1.TagListReq, res *fmtApiV1.TagListRes) *errcode.Error {
	err := svc.dao.TagPaginatedList(svc.ctx, req, res)
	if err != nil {
		return errcode.TagListFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) TagListAdmin(req *fmtAdminV1.TagListReq, res *fmtAdminV1.TagListRes) *errcode.Error {
	err := svc.dao.TagPaginatedListAdmin(svc.ctx, req, res)
	if err != nil {
		return errcode.TagListFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) TagAll(res *fmtApiV1.TagAllRes) *errcode.Error {
	err := svc.dao.TagAll(svc.ctx, res)
	if err != nil {
		return errcode.TagListFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) TagAllAdmin(res *fmtAdminV1.TagAllRes) *errcode.Error {
	err := svc.dao.TagAllAdmin(svc.ctx, res)
	if err != nil {
		return errcode.TagListFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) TagCreateAdmin(req *fmtAdminV1.TagCreateReq, res *fmtAdminV1.TagCreateRes) *errcode.Error {
	exists, err := svc.dao.TagExistsByTagName(svc.ctx, req.TagName, 0)
	if err != nil {
		return errcode.TagCreateFail.WithDetails(err.Error())
	}
	if exists {
		return errcode.TagCreateFail.WithDetails("标签名已存在")
	}
	tag := model.Tag{
		TagName:   req.TagName,
		TagStatus: req.TagStatus,
		Weight:    req.Weight,
	}
	err = svc.dao.TagCreate(svc.ctx, &tag)
	if err != nil {
		return errcode.TagCreateFail.WithDetails(err.Error())
	}
	res.Id = tag.Id
	return nil
}

func (svc *Service) TagUpdateAdmin(req *fmtAdminV1.TagUpdateReq, res *fmtAdminV1.TagUpdateRes) *errcode.Error {
	exists, err := svc.dao.UserExistsByUsername(svc.ctx, req.TagName, req.Id)
	if err != nil {
		return errcode.TagCreateFail.WithDetails(err.Error())
	}
	if exists {
		return errcode.TagCreateFail.WithDetails("标签名已存在")
	}
	err = svc.dao.TagUpdateAdmin(svc.ctx, req, res)
	if err != nil {
		return errcode.TagUpdateFail.WithDetails(err.Error())
	}
	return nil
}

func (svc *Service) TagDeleteAdmin(req *fmtAdminV1.TagDeleteReq, res *fmtAdminV1.TagDeleteRes) *errcode.Error {
	err := svc.dao.TagDelete(svc.ctx, req.Id)
	if err != nil {
		return errcode.TagDeleteFail.WithDetails(err.Error())
	}
	res.Id = req.Id
	return nil
}

func (svc *Service) TagDetail(id int, res *model.Tag) *errcode.Error {
	err := svc.dao.TagDetail(svc.ctx, id, res)
	if err != nil {
		return errcode.TagDetailFail.WithDetails(err.Error())
	}
	return nil
}
