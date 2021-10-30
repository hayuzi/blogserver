package service

import (
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func (svc *Service) TagList(req *fmtV1.TagListReq, res *fmtV1.TagListRes) *errcode.Error {
	err := svc.dao.TagPaginatedList(req, res)
	if err != nil {
		return errcode.TagListFail.WithDetails([]string{err.Error()}...)
	}
	return nil
}
