package dao

import (
	"context"
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
)

func (d *Dao) TagPaginatedList(ctx context.Context, req *fmtV1.TagListReq, res *fmtV1.TagListRes) error {
	res.Lists = make([]model.Tag, 0)
	pageNum, pageSize := app.InitPagination(req.PageNum, req.PageSize)
	offset := app.GetPageOffset(pageNum, pageSize)
	tx := d.engine.WithContext(ctx)
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

func (d *Dao) TagAll(ctx context.Context, res *fmtV1.TagAllRes) error {
	res.Lists = make([]model.Tag, 0)
	tx := d.engine.WithContext(ctx)
	err := tx.Model(&model.Tag{}).Order("weight DESC").Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) TagCreate(ctx context.Context, req *fmtV1.TagCreateReq, res *fmtV1.TagCreateRes) error {
	tag := model.Tag{
		TagName:   req.TagName,
		TagStatus: req.TagStatus,
		Weight:    req.Weight,
	}
	tx := d.engine.WithContext(ctx)
	if err := tx.Create(&tag).Error; err != nil {
		return err
	}
	res.Id = tag.Id
	return nil
}

func (d *Dao) TagUpdate(ctx context.Context, req *fmtV1.TagUpdateReq, res *fmtV1.TagUpdateRes) error {
	updateData := model.Tag{
		TagName:   req.TagName,
		TagStatus: req.TagStatus,
		Weight:    req.Weight,
	}
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", req.Id).Updates(&updateData).Error; err != nil {
		return err
	}
	res.Id = req.Id
	return nil
}

func (d *Dao) TagDelete(ctx context.Context, req *fmtV1.TagDeleteReq, res *fmtV1.TagDeleteRes) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", req.Id).Delete(&model.Tag{}).Error; err != nil {
		return err
	}
	res.Id = req.Id
	return nil
}

func (d *Dao) TagDetail(ctx context.Context, id int, res *model.Tag) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", id).Delete(&res).Error; err != nil {
		return err
	}
	res.Id = id
	return nil
}
