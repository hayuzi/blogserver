package dao

import (
	"context"
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
)

func (d *Dao) TagCount(tagName string, tagStatus int) (int64, error) {
	var count int64
	tx := d.engine.WithContext(context.Background())
	if tagName != "" {
		tx = tx.Where("tag_name = ?", tagName)
	}
	if tagStatus > 0 {
		tx = tx.Where("tag_status = ?", tagStatus)
	}
	err := tx.Model(&model.Tag{}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (d *Dao) TagPaginatedList(req *fmtV1.TagListReq, res *fmtV1.TagListRes) error {
	res.Lists = make([]model.Tag, 0)
	pageNum, pageSize := app.InitPagination(req.PageNum, req.PageSize)
	offset := app.GetPageOffset(pageNum, pageSize)
	tx := d.engine.WithContext(context.Background())
	if req.TagName != "" {
		tx = tx.Where("tag_name = ?", req.TagName)
	}
	if req.TagStatus > 0 {
		tx = tx.Where("tag_status = ?", req.TagStatus)
	}
	err := tx.Model(&model.Tag{}).Count(&res.Total).Error
	if err != nil {
		return err
	}
	err = tx.Model(&model.Tag{}).Offset(offset).Limit(pageSize).Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}
